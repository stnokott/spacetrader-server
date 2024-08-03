package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/convert"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
)

var buildSystemIndexTimeout = 10 * time.Minute

// UpdateSystemIndex queries all systems from the API and writes them to the DB.
// This index can be used later to query systems quickly without relying on the API.
// This approach is valid since systems are expected to be static.
//
// This function is blocking.
func (s *Server) UpdateSystemIndex(force bool) error {
	log.WithField("timeout", buildSystemIndexTimeout).Info("building system index")

	ctx, cancel := context.WithTimeout(context.Background(), buildSystemIndexTimeout)
	defer cancel()

	if !force {
		hasIndex, err := s.hasSystems(ctx)
		if err != nil {
			return fmt.Errorf("checking for system index: %w", err)
		}
		if hasIndex {
			log.Info("system index exists, skipping refresh")
			return nil
		}
	} else {
		log.Info("forcing system index refresh")
	}

	return s.replaceSystems(ctx)
}

// replaceSystems replaces the contents of the `systems` table with results from systemChan.
// It continues reading from systemChan until it is closed or ctx expires.
func (s *Server) replaceSystems(ctx context.Context) (err error) {
	var systems []*api.System
	systems, err = getPaginated[*api.System](
		ctx,
		s,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/systems?page=%d&limit=20", page)
		},
	)
	if err != nil {
		err = fmt.Errorf("querying systems: %w", err)
		return
	}

	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		err = fmt.Errorf("creating SQLite transaction: %w", err)
		return
	}
	defer func() {
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				log.Errorf("failed to rollback: %v", errRollback)
			}
		} else {
			if errCommit := tx.Commit(); errCommit != nil {
				log.Errorf("failed to commit: %v", errCommit)
			}
		}
	}()

	q := s.query.WithTx(tx)

	// delete existing index
	if err = q.TruncateSystems(ctx); err != nil {
		return
	}

	defer func() {
		if err == nil {
			log.WithField("n", len(systems)).Info("system index replaced")
		}
	}()

	for i, system := range systems {
		if _, contextExceeded := <-ctx.Done(); contextExceeded {
			err = fmt.Errorf("context exceeded after %d systems processed", i)
			return
		}
		factions := make([]string, len(system.Factions))
		for j, fac := range system.Factions {
			factions[j] = string(fac.Symbol)
		}

		if err = s.query.InsertSystem(ctx, query.InsertSystemParams{
			Symbol:   system.Symbol,
			X:        int64(system.X),
			Y:        int64(system.Y),
			Type:     string(system.Type),
			Factions: strings.Join(factions, ","),
		}); err != nil {
			err = fmt.Errorf("inserting system '%s': %v", system.Symbol, err)
			return
		}
	}
	return
}

// hasSystems returns true if the systems table has at least one row, indicating
// an existing Systems index.
func (s *Server) hasSystems(ctx context.Context) (bool, error) {
	result, err := s.db.QueryContext(ctx, "SELECT 1 FROM systems LIMIT 1")
	if err != nil {
		return false, err
	}
	hasNext := result.Next()
	_ = result.Close()
	return hasNext, nil
}

// GetSystemsInRect streams all systems whose coordinates fall into rect.
func (s *Server) GetSystemsInRect(rect *pb.Rect, stream pb.Spacetrader_GetSystemsInRectServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.query.SelectSystemsInRect(ctx, query.SelectSystemsInRectParams{
		XMin: int64(rect.Start.X),
		YMin: int64(rect.Start.Y),
		XMax: int64(rect.End.X),
		YMax: int64(rect.End.Y),
	})
	if err != nil {
		return fmt.Errorf("querying systems within rect: %w", err)
	}

	for _, row := range rows {
		system, err := convert.ConvertSystem(&row)
		if err != nil {
			return err
		}

		if err = stream.Send(system); err != nil {
			return fmt.Errorf("sending system via gRPC: %w", err)
		}
	}
	return nil
}
