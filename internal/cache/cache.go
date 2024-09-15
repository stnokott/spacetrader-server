// Package cache implements several caches for API data.
package cache

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/convert"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	"github.com/stnokott/spacetrader-server/internal/log"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
)

var logger = log.ForComponent("cache")

// TODO: system cache should store the time at which it created the cache.
//       this should then be used to check on each startup whether the game was reset
//       since last cache creation, thus requiring recreation of the cache.

// SystemCache is a cache for galaxy systems.
type SystemCache struct {
	client  *api.Client
	db      *sql.DB
	queries *query.Queries
}

func NewSystemCache(client *api.Client, db *sql.DB, queries *query.Queries) SystemCache {
	return SystemCache{
		client:  client,
		db:      db,
		queries: queries,
	}
}

// Create populates the contents of the `systems` table with results from the API.
func (c SystemCache) Create(ctx context.Context, progressChan chan<- float64) error {
	// check if already exists
	if v, err := c.queries.HasSystemsRows(ctx); err != nil {
		return err
	} else if v != 0 {
		return nil
	}

	tx, err := query.WithTx(ctx, c.db, c.queries)
	if err != nil {
		return err
	}

	err = c.populateWithTx(ctx, tx, progressChan)
	return tx.Done(err)
}

func (c SystemCache) populateWithTx(ctx context.Context, tx query.Tx, progressChan chan<- float64) error {
	systemsIter := api.GetPaginated[*api.System](
		ctx,
		c.client,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/systems?page=%d&limit=20", page)
		},
	)

	// delete existing index
	logger.Debug("clearing existing system/waypoint/jumpgate index")
	if err := errors.Join(tx.TruncateSystems(ctx), tx.TruncateSystems(ctx), tx.TruncateJumpGates(ctx)); err != nil {
		return fmt.Errorf("truncating system/waypoint/jumpgate index: %w", err)
	}

	total := 0
	n := 0

	for systemPage, errPage := range systemsIter {
		if errPage != nil {
			return fmt.Errorf("querying systems: %w", errPage)
		}
		if err := c.insertSystemPage(ctx, tx, systemPage.Items); err != nil {
			return err
		}
		total = systemPage.Total
		n += len(systemPage.Items)
		progressChan <- float64(n) / float64(total)
	}
	return nil
}

func (c SystemCache) insertSystemPage(ctx context.Context, tx query.Tx, page []*api.System) error {
	for _, system := range page {
		factions := make([]string, len(system.Factions))
		for i, fac := range system.Factions {
			factions[i] = string(fac.Symbol)
		}

		// TODO: use converter
		if err := tx.InsertSystem(ctx, query.InsertSystemParams{
			Symbol:   system.Symbol,
			X:        int64(system.X),
			Y:        int64(system.Y),
			Type:     string(system.Type),
			Factions: strings.Join(factions, ","),
		}); err != nil {
			return fmt.Errorf("inserting system '%s': %w", system.Symbol, err)
		}

		for _, wp := range system.Waypoints {
			if err := c.insertWaypoint(ctx, system.Symbol, &wp, tx); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c SystemCache) insertWaypoint(ctx context.Context, system string, wp *api.SystemWaypoint, tx query.Tx) error {
	if err := tx.InsertWaypoint(ctx, query.InsertWaypointParams{
		Symbol: wp.Symbol,
		System: system,
		X:      int64(wp.X),
		Y:      int64(wp.Y),
		Orbits: wp.Orbits,
		Type:   string(wp.Type),
	}); err != nil {
		return fmt.Errorf("inserting waypoint '%s': %w", wp.Symbol, err)
	}

	if wp.Type == api.JUMPGATE {
		if err := c.populateJumpgateWaypoint(ctx, system, wp.Symbol, tx); err != nil {
			return err
		}
	}

	return nil
}

func (c SystemCache) populateJumpgateWaypoint(ctx context.Context, system string, wp string, tx query.Tx) error {
	// check if waypoint if charted (because if it isn't, we dont have jumpgate info)
	// also, this information isn't available in the SystemWaypoint type, so we need to waste an API call for checking.
	url := fmt.Sprintf("/systems/%s/waypoints/%s", system, wp)
	waypoint := &struct {
		Data api.Waypoint `json:"data"`
	}{}
	if err := c.client.Get(ctx, waypoint, url); err != nil {
		return fmt.Errorf("querying waypoint: %w", err)
	}
	if waypoint.Data.Chart == nil {
		// not charted => no jumpgate info => abort
		return nil
	}

	url = fmt.Sprintf("/systems/%s/waypoints/%s/jump-gate", system, wp)
	jump := &struct {
		Data api.JumpGate `json:"data"`
	}{}
	if err := c.client.Get(ctx, jump, url); err != nil {
		return fmt.Errorf("querying jump gate: %w", err)
	}

	for _, connection := range jump.Data.Connections {
		if err := tx.InsertJumpGate(ctx, query.InsertJumpGateParams{
			Waypoint:   wp,
			ConnectsTo: connection,
		}); err != nil {
			return fmt.Errorf("inserting jumpgate: %w", err)
		}
	}
	return nil
}

// FleetCache is an in-memory cache of all player-owned ships.
type FleetCache struct {
	Ships  []*pb.Ship
	client *api.Client
}

func NewFleetCache(client *api.Client) *FleetCache {
	return &FleetCache{
		client: client,
	}
}

// Create (re)populates the cache from the API.
func (c *FleetCache) Create(ctx context.Context, progressChan chan<- float64) error {
	progressChan <- 0
	shipsIter := api.GetPaginated[*api.Ship](
		ctx,
		c.client,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/my/ships?page=%d&limit=20", page)
		},
	)

	ships, err := api.CollectPages(shipsIter)
	if err != nil {
		return fmt.Errorf("querying ships: %w", err)
	}

	c.Ships, err = convert.ConvertShips(ships)
	if err != nil {
		return err
	}
	return nil
}

/*
// ShipByName returns a ship from the cache by its name.
//
// An error is returned when no ship with that name is found.
func (c *FleetCache) ShipByName(name string) (*pb.Ship, error) {
	for _, ship := range c.Ships {
		if ship.Id == name {
			return ship, nil
		}
	}
	return nil, fmt.Errorf("no ship with name '%s' found", name)
}
*/
