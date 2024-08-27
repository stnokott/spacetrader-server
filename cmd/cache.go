package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/convert"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
)

// TODO: write functions for querying API stuff (e.g. getSystem)
// which wrap API calls, but also handle caching.
// So when calling getSystem and the queried system doesn't exist in the cache yet,
// we query the API and write the result to the cache.

type cacheManager struct {
	caches map[string]cache
}

type progressFunc func(total int, current int)

type cache interface {
	Create(ctx context.Context, srv *Server, progressFunc progressFunc) error
}

func (cm cacheManager) Create(ctx context.Context, srv *Server) error {
	g, ctxCache := errgroup.WithContext(ctx)

	type progress [2]int // total, current
	var cacheProgress sync.Map

	// start cache workers
	for name, cache := range cm.caches {
		g.Go(func() error {
			return cache.Create(ctxCache, srv, func(total int, current int) {
				cacheProgress.Store(name, progress{total, current})
			})
		})
	}

	// channel for stopping the consumer
	stopChan := make(chan struct{})
	defer func() {
		stopChan <- struct{}{}
	}()

	// start progress consumer
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-stopChan:
				return
			case <-ticker.C:
				cacheProgress.Range(func(key, value any) bool {
					v := value.(progress)
					log.Infof("Cache %s at %d/%d", key, v[1], v[0])
					return true
				})
			}
		}
	}()

	return g.Wait()
}

var indexTimeout = 1 * time.Hour

// CreateCaches updates or creates all registered indexes.
// It should be called once at the beginning of the program loop.
func (s *Server) CreateCaches(ctxParent context.Context) error {
	log.Info("creating caches")

	ctx, cancel := context.WithTimeout(ctxParent, indexTimeout)
	defer cancel()

	cacheManager := cacheManager{
		caches: map[string]cache{
			"Systems": s.systemCache,
			"Fleet":   s.fleetCache,
		},
	}

	return cacheManager.Create(ctx, s)
}

// SystemCache is a cache for galaxy systems.
type SystemCache struct{}

var _ cache = SystemCache{}

// Create populates the contents of the `systems` table with results from the API.
func (c SystemCache) Create(ctx context.Context, srv *Server, progressFunc progressFunc) error {
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

	err = c.populateWithTx(ctx, srv, tx, progressFunc)
	return errors.Join(err, tx.Done(err))
}

func (c SystemCache) populateWithTx(ctx context.Context, srv *Server, tx query.Tx, progressFunc progressFunc) error {
	systemsIter := getPaginated[*api.System](
		ctx,
		srv,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/systems?page=%d&limit=20", page)
		},
	)

	// delete existing index
	log.Debug("clearing existing system/waypoint/jumpgate index")
	if err := errors.Join(tx.TruncateSystems(ctx), tx.TruncateSystems(ctx), tx.TruncateJumpGates(ctx)); err != nil {
		return fmt.Errorf("truncating system/waypoint/jumpgate index: %w", err)
	}

	total := 0
	n := 0

	for systemPage, errPage := range systemsIter {
		if errPage != nil {
			return fmt.Errorf("querying systems: %w", errPage)
		}
		if err := c.insertSystemPage(ctx, srv, tx, systemPage.Items); err != nil {
			return err
		}
		total = systemPage.Total
		n += len(systemPage.Items)
		progressFunc(total, n)
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

		for _, wp := range system.Waypoints {
			if err := c.insertWaypoint(ctx, system.Symbol, &wp, srv, tx); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c SystemCache) insertWaypoint(ctx context.Context, system string, wp *api.SystemWaypoint, srv *Server, tx query.Tx) error {
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
		if err := c.populateJumpgateWaypoint(ctx, system, wp.Symbol, srv, tx); err != nil {
			return err
		}
	}

	return nil
}

func (SystemCache) populateJumpgateWaypoint(ctx context.Context, system string, wp string, srv *Server, tx query.Tx) error {
	// check if waypoint if charted (because if it isn't, we dont have jumpgate info)
	// also, this information isn't available in the SystemWaypoint type, so we need to waste an API call for checking.
	url := fmt.Sprintf("/systems/%s/waypoints/%s", system, wp)
	waypoint := &struct {
		Data api.Waypoint `json:"data"`
	}{}
	if err := srv.get(ctx, waypoint, url); err != nil {
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
	if err := srv.get(ctx, jump, url); err != nil {
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
	Ships []*pb.Ship
}

var _ cache = (*FleetCache)(nil)

// Create (re)populates the cache from the API.
func (c *FleetCache) Create(ctx context.Context, srv *Server, progressFunc progressFunc) error {
	progressFunc(1, 0)
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

	if c.Ships, err = convert.ConvertShips(ships); err != nil {
		return fmt.Errorf("converting ship: %w", err)
	}
	progressFunc(1, 1)
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
