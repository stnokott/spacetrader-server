package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/convert"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
)

// Cache stores API-related data locally for faster access.
type Cache interface {
	// Exists returns false if the cache has not been created and otherwise false.
	Exists(ctx context.Context, srv *Server) (bool, error)
	// Update creates/refreshes the cache.
	Update(ctx context.Context, srv *Server) error
}

var indexTimeout = 5 * time.Minute

// CreateIndexes updates or creates all registered indexes.
// It should be called once at the beginning of the program loop.
func (s *Server) CreateIndexes(ctxParent context.Context) (err error) {
	ctx, cancel := context.WithTimeout(ctxParent, indexTimeout)
	defer cancel()

	log.Info("updating indexes")

	var exists bool
	caches := map[string]Cache{
		"Systems": s.systemCache,
		"Fleet":   s.fleetCache,
	}
	for name, cache := range caches {
		exists, err = cache.Exists(ctx, s)
		if err != nil {
			return
		}
		if exists {
			log.WithField("cache_name", name).Debug("cache already initialized")
			continue
		}
		log.WithField("cache_name", name).Debug("cache requires initialization")
		if err = cache.Update(ctx, s); err != nil {
			return
		}
	}
	return
}

// DBCache implements the Cache interface.
// It provides wrappers around database-related operations, like transactions.
type DBCache struct {
	existsFunc func(ctx context.Context, q *query.Queries) (bool, error)
	updateFunc func(ctx context.Context, srv *Server, tx query.Tx) error
}

var _ Cache = (*DBCache)(nil)

// NewDBCache creates a new DBCache, implementing the Cache interface.
func NewDBCache(
	existsFunc func(ctx context.Context, q *query.Queries) (bool, error),
	updateFunc func(ctx context.Context, srv *Server, tx query.Tx) error,
) Cache {
	return &DBCache{
		existsFunc: existsFunc,
		updateFunc: updateFunc,
	}
}

// Exists implements Cache.
func (c *DBCache) Exists(ctx context.Context, srv *Server) (bool, error) {
	return c.existsFunc(ctx, srv.query)
}

// Update implements Cache.
func (c *DBCache) Update(ctx context.Context, srv *Server) error {
	tx, err := query.WithTx(ctx, srv.db, srv.query)
	if err != nil {
		return err
	}
	err = c.updateFunc(ctx, srv, tx)
	return errors.Join(err, tx.Done(err))
}

// NewSystemCache creates a new cache for systems.
func NewSystemCache() Cache {
	return NewDBCache(systemCacheExists, updateSystemCache)
}

// systemCacheExists returns true if the systems table has at least one row, indicating
// an existing Systems index.
func systemCacheExists(ctx context.Context, q *query.Queries) (bool, error) {
	v, err := q.HasSystemsRows(ctx)
	return v != 0, err
}

// updateSystemCache replaces the contents of the `systems` table with results from the API.
// It continues reading from systemChan until it is closed or ctx expires.
func updateSystemCache(ctx context.Context, srv *Server, tx query.Tx) error {
	systemsIter := getPaginated[*api.System](
		ctx,
		srv,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/systems?page=%d&limit=20", page)
		},
	)

	// delete existing index
	log.Debug("clearing existing system index")
	if err := tx.TruncateSystems(ctx); err != nil {
		return fmt.Errorf("truncating systems index: %w", err)
	}

	for systemPage, errPage := range systemsIter {
		if errPage != nil {
			return fmt.Errorf("querying systems: %w", errPage)
		}
		if err := insertSystemPage(ctx, tx, systemPage); err != nil {
			return err
		}
	}
	return nil
}

func insertSystemPage(ctx context.Context, tx query.Tx, page []*api.System) error {
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
			return fmt.Errorf("inserting system '%s': %v", system.Symbol, err)
		}
	}

	return nil
}

// FleetCache is an in-memory cache of all player-owned ships.
type FleetCache struct {
	ships []*pb.Ship
}

var _ Cache = (*FleetCache)(nil)

// Exists implements Cache.
func (c *FleetCache) Exists(_ context.Context, _ *Server) (bool, error) {
	return c.ships != nil, nil
}

// Update implements Cache.
func (c *FleetCache) Update(ctx context.Context, srv *Server) error {
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
	c.ships = converted
	return nil
}

// ShipByName returns a ship from the cache by its name.
//
// An error is returned when no ship with that name is found.
func (c *FleetCache) ShipByName(name string) (*pb.Ship, error) {
	for _, ship := range c.ships {
		if ship.Id == name {
			return ship, nil
		}
	}
	return nil, fmt.Errorf("no ship with name '%s' found", name)
}
