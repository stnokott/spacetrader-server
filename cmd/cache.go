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

var indexTimeout = 5 * time.Minute

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

func (SystemCache) createWithTx(ctx context.Context, srv *Server, tx query.Tx) error {
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
	Ships []*pb.Ship
}

// Create (re)populates the cache from the API.
func (c FleetCache) Create(ctx context.Context, srv *Server) error {
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
func (c FleetCache) ShipByName(name string) (*pb.Ship, error) {
	for _, ship := range c.Ships {
		if ship.Id == name {
			return ship, nil
		}
	}
	return nil, fmt.Errorf("no ship with name '%s' found", name)
}
