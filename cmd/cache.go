package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/convert"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
)

var buildSystemIndexTimeout = 20 * time.Minute

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

// replaceSystems replaces the contents of the `systems` table with results from the API.
// It continues reading from systemChan until it is closed or ctx expires.
func (s *Server) replaceSystems(ctx context.Context) (err error) {
	log.Info("step 1/2: querying systems from API")
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

	log.Infof("step 2/2: inserting %d systems into DB", len(systems))

	tx, err := query.WithTx(ctx, s.db, s.query)
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Done(err)
	}()

	// delete existing index
	log.Debug("clearing existing system index")
	if err = tx.TruncateSystems(ctx); err != nil {
		return
	}

	defer func() {
		if err == nil {
			log.WithField("n", len(systems)).Info("system index replaced")
		}
	}()

	for i, system := range systems {
		select {
		case <-ctx.Done():
			err = fmt.Errorf("context exceeded after %d systems processed", i)
			return
		default:
		}
		factions := make([]string, len(system.Factions))
		for j, fac := range system.Factions {
			factions[j] = string(fac.Symbol)
		}

		if err = tx.InsertSystem(ctx, query.InsertSystemParams{
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

// GetFleet returns the complete list of ships in the agent's posession.
func (s *Server) GetFleet(ctx context.Context, _ *emptypb.Empty) (fleet *pb.Fleet, err error) {
	var ships []*api.Ship
	ships, err = getPaginated[*api.Ship](
		ctx,
		s,
		func(page int) (urlPath string) {
			return fmt.Sprintf("/my/ships?page=%d&limit=20", page)
		},
	)
	if err != nil {
		return nil, fmt.Errorf("querying ships: %w", err)
	}

	converted := make([]*pb.Ship, len(ships))

	tx, err := query.WithTx(ctx, s.db, s.query)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = tx.Done(err)
	}()

	if err = tx.TruncateShips(ctx); err != nil {
		return nil, fmt.Errorf("truncating ships table: %w", err)
	}

	for i, ship := range ships {
		if err := tx.InsertShip(ctx, query.InsertShipParams{
			Symbol:          ship.Symbol,
			CurrentSystem:   ship.Nav.SystemSymbol,
			CurrentWaypoint: ship.Nav.WaypointSymbol,
		}); err != nil {
			return nil, fmt.Errorf("inserting ship into DB: %w", err)
		}

		if converted[i], err = convert.ConvertShip(ship); err != nil {
			return nil, fmt.Errorf("converting ship: %w", err)
		}
	}

	fleet = &pb.Fleet{Ships: converted}
	return
}

// GetShipCoordinates returns the x and y coordinates for a ship, identified by its name
func (s *Server) GetShipCoordinates(ctx context.Context, req *pb.GetShipCoordinatesRequest) (*pb.GetShipCoordinatesResponse, error) {
	result := new(struct {
		Ship *api.Ship `json:"data"`
	})
	if err := s.get(ctx, result, "/my/ships/"+req.ShipName, 200); err != nil {
		return nil, err
	}

	if result.Ship == nil {
		return nil, fmt.Errorf("no ship '%s' found in fleet", req.ShipName)
	}
	systemName := result.Ship.Nav.SystemSymbol
	system, err := s.query.GetSystemByName(ctx, systemName)
	if err != nil {
		return nil, err
	}
	return &pb.GetShipCoordinatesResponse{
		X: int32(system.X), Y: int32(system.Y),
	}, nil
}

// GetSystemsInRect streams all systems whose coordinates fall into rect.
func (s *Server) GetSystemsInRect(rect *pb.Rect, stream pb.Spacetrader_GetSystemsInRectServer) error {
	ctx, cancel := context.WithTimeout(stream.Context(), 5*time.Second)
	defer cancel()

	rows, err := s.query.GetSystemsInRect(ctx, query.GetSystemsInRectParams{
		XMin: int64(rect.Start.X),
		YMin: int64(rect.Start.Y),
		XMax: int64(rect.End.X),
		YMax: int64(rect.End.Y),
	})
	if err != nil {
		return fmt.Errorf("querying systems within rect: %w", err)
	}

	for _, row := range rows {
		system, err := convert.ConvertSystem(&row.System)
		if err != nil {
			return err
		}

		if err = stream.Send(&pb.GetSystemsInRectResponse{
			System:    system,
			ShipCount: int32(row.ShipCount),
		}); err != nil {
			return fmt.Errorf("sending system via gRPC: %w", err)
		}
	}
	return nil
}
