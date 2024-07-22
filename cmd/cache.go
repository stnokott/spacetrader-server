package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stnokott/spacetrader/internal/api"
	pb "github.com/stnokott/spacetrader/internal/proto"
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
	dataChan, stopChan := getPaginatedAsync[*api.System](
		ctx,
		s,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/systems?page=%d&limit=20", page)
		},
	)
	defer func() {
		// stop querying when er encounter an error
		if err != nil {
			stopChan <- struct{}{}
		}
	}()

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

	// delete existing index
	if _, err = tx.ExecContext(ctx, "DELETE FROM systems"); err != nil {
		return
	}

	// create prepared statement
	var insert *sql.Stmt
	insert, err = tx.PrepareContext(ctx, `
		INSERT INTO systems VALUES (
			$symbol, $x, $y, $type, $factions
		);
	`)
	if err != nil {
		err = fmt.Errorf("preparing INSERT statement: %w", err)
		return
	}

	n := 0
	defer func() {
		if err == nil {
			log.WithField("n", n).Info("system index replaced")
		}
	}()

	for {
		select {
		case <-ctx.Done():
			err = fmt.Errorf("context exceeded after %d systems processed", n)
			return
		case rcv, ok := <-dataChan:
			if !ok {
				return
			}
			if rcv.Err != nil {
				err = rcv.Err
				return
			}
			system := rcv.Data

			factions := make([]string, len(system.Factions))
			for i, fac := range system.Factions {
				factions[i] = string(fac.Symbol)
			}
			if _, err = insert.ExecContext(
				ctx,
				sql.Named("symbol", system.Symbol),
				sql.Named("x", system.X),
				sql.Named("y", system.Y),
				sql.Named("type", string(system.Type)),
				sql.Named("factions", strings.Join(factions, ",")),
			); err != nil {
				err = fmt.Errorf("inserting system '%s': %v", system.Symbol, err)
				return
			}
			n++
		}
	}
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

const _sqlGetSystemsInRect = `
	SELECT symbol, x, y, type, factions FROM systems
	WHERE TRUE
		AND x >= $x_min AND x <= $x_max
		AND y >= $y_min AND y <= $y_max
`

// GetSystemsInRect streams all systems whose coordinates fall into rect.
func (s *Server) GetSystemsInRect(rect *pb.Rect, stream pb.Spacetrader_GetSystemsInRectServer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(
		ctx,
		_sqlGetSystemsInRect,
		sql.Named("x_min", rect.Start.X),
		sql.Named("y_min", rect.Start.Y),
		sql.Named("x_max", rect.End.X),
		sql.Named("y_max", rect.End.Y),
	)
	if err != nil {
		return fmt.Errorf("querying systems within rect: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	type result struct {
		Symbol   string
		X        int32
		Y        int32
		Type     string
		Factions string
	}
	for rows.Next() {
		dst := result{}
		if err = rows.Scan(&dst.Symbol, &dst.X, &dst.Y, &dst.Type, &dst.Factions); err != nil {
			return fmt.Errorf("reading system from query result: %w", err)
		}
		if err = stream.Send(&pb.System{
			Id:       dst.Symbol,
			X:        dst.X,
			Y:        dst.Y,
			Type:     pb.System_BLACK_HOLE,           // TODO: use correct type
			Factions: []pb.Faction{pb.Faction_AEGIS}, // TODO: use correct type
		}); err != nil {
			return fmt.Errorf("sending system via gRPC: %w", err)
		}
	}
	return nil
}
