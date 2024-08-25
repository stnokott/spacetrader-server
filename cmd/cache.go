package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/convert"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
)

// TODO: for waypoint cache, implement on-demand caching:
// perform waypoint caching in the background, slowly traversing all known systems
// when a waypoint is queried that is not cached yet, prioritize that waypoin

var indexTimeout = 3 * time.Hour

// CreateCaches updates or creates all registered indexes.
// It should be called once at the beginning of the program loop.
func (s *Server) CreateCaches(ctxParent context.Context) (err error) {
	log.Info("creating caches")

	ctx, cancel := context.WithTimeout(ctxParent, indexTimeout)
	defer cancel()

	var g errgroup.Group

	g.Go(func() error {
		return s.systemCache.Create(ctx, s)
	})
	g.Go(func() error {
		return s.fleetCache.Create(ctx, s)
	})
	if err = g.Wait(); err != nil {
		err = fmt.Errorf("creating caches: %w", err)
		return
	}
	log.Info("caches created")
	return
}

// SystemCache is a cache for galaxy systems.
type SystemCache struct{}

func (c SystemCache) createWithTx(ctx context.Context, srv *Server, tx query.Tx) error {
	systemsIter := getPaginated[*api.System](
		ctx,
		srv,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/systems?page=%d&limit=20", page)
		},
	)

	// delete existing index
	log.Debug("clearing existing system/waypoint index")
	if err := errors.Join(tx.TruncateSystems(ctx), tx.TruncateSystems(ctx)); err != nil {
		return fmt.Errorf("truncating system/waypoint index: %w", err)
	}

	for systemPage, errPage := range systemsIter {
		if errPage != nil {
			return fmt.Errorf("querying systems: %w", errPage)
		}
		if err := c.insertSystemPage(ctx, srv, tx, systemPage); err != nil {
			return err
		}
	}
	return nil
}

func (c SystemCache) insertSystemPage(ctx context.Context, srv *Server, tx query.Tx, page []*api.System) error {
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

		if err := c.createWaypointsForSystem(ctx, system.Symbol, srv, tx); err != nil {
			return err
		}
	}
	return nil
}

func (c SystemCache) createWaypointsForSystem(ctx context.Context, system string, srv *Server, tx query.Tx) error {
	waypointsIter := getPaginated[*api.Waypoint](
		ctx,
		srv,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/systems/%s/waypoints?page=%d&limit=20", system, page)
		},
	)

	for waypointPage, errPage := range waypointsIter {
		if errPage != nil {
			return fmt.Errorf("querying waypoints: %w", errPage)
		}
		if err := c.insertWaypointPage(ctx, tx, waypointPage); err != nil {
			return err
		}
	}
	return nil
}

func (SystemCache) insertWaypointPage(ctx context.Context, tx query.Tx, page []*api.Waypoint) error {
	for _, wp := range page {
		if err := tx.InsertWaypoint(ctx, query.InsertWaypointParams{
			Symbol: wp.Symbol,
			System: wp.SystemSymbol,
			Orbits: wp.Orbits,
			X:      int64(wp.X),
			Y:      int64(wp.Y),
			Type:   string(wp.Type),
		}); err != nil {
			return fmt.Errorf("inserting waypoint '%s': %w", wp.Symbol, err)
		}
	}
	return nil
}

// Create populates the contents of the `systems` table with results from the API.
func (c SystemCache) Create(ctx context.Context, srv *Server) error {
	// check if already exists
	if v, err := srv.query.HasSystemsRows(ctx); err != nil {
		return err
	} else if v != 0 {
		return nil
	}

	tx, err := query.WithTx(ctx, srv.db, srv.query)
	if err != nil {
		return err
	}
	err = c.createWithTx(ctx, srv, tx)
	return errors.Join(err, tx.Done(err))
}

// FleetCache is an in-memory cache of all player-owned ships.
type FleetCache struct {
	Ships []*pb.Ship
}

// Create (re)populates the cache from the API.
func (c *FleetCache) Create(ctx context.Context, srv *Server) error {
	shipsIter := getPaginated[*api.Ship](
		ctx,
		srv,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/my/ships?page=%d&limit=20", page)
		},
	)

	ships, err := collectPages(shipsIter)
	if err != nil {
		return fmt.Errorf("querying ships: %w", err)
	}

	converted := make([]*pb.Ship, len(ships))
	for i, ship := range ships {
		if converted[i], err = convert.ConvertShip(ship); err != nil {
			return fmt.Errorf("converting ship: %w", err)
		}
	}
	c.Ships = converted
	return nil
}

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
