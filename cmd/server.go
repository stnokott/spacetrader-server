package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/cache"
	"github.com/stnokott/spacetrader-server/internal/convert"
	"github.com/stnokott/spacetrader-server/internal/log"
	"github.com/stnokott/spacetrader-server/internal/worker"
	"google.golang.org/grpc"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/stnokott/spacetrader-server/internal/db/query"
	pb "github.com/stnokott/spacetrader-server/internal/proto"
)

var logger = log.ForComponent("server")

// Server performs requests to the SpaceTraders API and offers them via gRPC.
type Server struct {
	api     *api.Client
	db      *sql.DB
	queries *query.Queries
	// TODO: check if api, db or query can be removed

	systemCache cache.SystemCache
	fleetCache  *cache.FleetCache

	worker *worker.Worker

	pb.UnimplementedSpacetraderServer
}

// New creates and returns a new Client instance.
func New(baseURL string, token string, dbFile string) (*Server, error) {
	client := api.NewClient(baseURL, token)

	db, err := newDB(dbFile)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	q, err := query.Prepare(ctx, db)
	if err != nil {
		return nil, err
	}

	worker := worker.NewWorker()
	worker.Start(context.TODO(), true)

	return &Server{
		api:     client,
		db:      db,
		queries: q,

		systemCache: cache.NewSystemCache(client, db, q),
		fleetCache:  cache.NewFleetCache(client),
		worker:      worker,
	}, nil
}

// Close terminates all underlying connections.
func (s *Server) Close() error {
	return errors.Join(s.queries.Close(), s.db.Close())
}

// Listen starts the gRPC server on the specified port.
//
// It is blocking.
func (s *Server) Listen(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("TCP listen: %w", err)
	}
	srv := grpc.NewServer(
		grpc.ChainStreamInterceptor(onGrpcStream),
		grpc.ChainUnaryInterceptor(onGrpcUnary),
	)
	pb.RegisterSpacetraderServer(srv, s)
	logger.Infof("gRPC server listening on port %d", port)
	if err := srv.Serve(lis); err != nil {
		return fmt.Errorf("TCP serve: %w", err)
	}
	return nil
}

func onGrpcStream(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logger.Debugf("gRPC call to %s", info.FullMethod)
	err := handler(srv, stream)
	if err != nil {
		logger.Errorf("gRPC error streaming %s: %v", info.FullMethod, err)
	}
	return err
}

func onGrpcUnary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	logger.Debugf("gRPC call to %s", info.FullMethod)
	resp, err = handler(ctx, req)
	if err != nil {
		logger.Errorf("gRPC error unary %s: %v", info.FullMethod, err)
	}
	return
}

// TODO: write functions for querying API stuff (e.g. getSystem)
// which wrap API calls, but also handle caching.
// So when calling getSystem and the queried system doesn't exist in the cache yet,
// we query the API and write the result to the cache.

var indexTimeout = 1 * time.Hour

// CreateCaches updates or creates all registered indexes.
// It should be called once at the beginning of the program loop.
func (s *Server) CreateCaches(ctxParent context.Context) error {
	logger.Info("creating caches")

	ctx, cancel := context.WithTimeout(ctxParent, indexTimeout)
	defer cancel()

	err := s.worker.AddAndWait(ctx, "create-system-cache", func(ctx context.Context, progressChan chan<- float64) error {
		return s.systemCache.Create(ctx, progressChan)
	})
	if err != nil {
		return err
	}
	err = s.worker.AddAndWait(ctx, "create-fleet-cache", func(ctx context.Context, progressChan chan<- float64) error {
		return s.fleetCache.Create(ctx, progressChan)
	})
	if err != nil {
		return err
	}

	return nil
}

// Ping is used by clients to ensure this server is online.
func (s *Server) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// GetServerStatus returns the current server status and some statistics.
func (s *Server) GetServerStatus(ctx context.Context, _ *emptypb.Empty) (*pb.ServerStatus, error) {
	result := new(api.Status)
	if err := s.api.Get(ctx, result, "/"); err != nil {
		return nil, err
	}

	return convert.ConvertStatus(result), nil
}

// GetCurrentAgent returns information about the agent identified by the current token.
func (s *Server) GetCurrentAgent(ctx context.Context, _ *emptypb.Empty) (*pb.Agent, error) {
	result := new(struct {
		// for some reason, SpaceTraders decided it's a good idea to wrap the agent
		// info in a useless "data" field.
		Data api.Agent `json:"data"`
	})
	if err := s.api.Get(ctx, result, "/my/agent"); err != nil {
		return nil, err
	}

	return convert.ConvertAgent(&result.Data)
}

// GetFleet returns the complete list of ships in the agent's posession.
func (s *Server) GetFleet(_ context.Context, _ *emptypb.Empty) (*pb.Fleet, error) {
	if s.fleetCache.Ships == nil {
		return nil, errors.New("fleet cache has not been initialized")
	}
	return &pb.Fleet{
		Ships: s.fleetCache.Ships,
	}, nil
}

// GetShipCoordinates returns the x and y coordinates for a ship, identified by its name
func (s *Server) GetShipCoordinates(ctx context.Context, req *pb.GetShipCoordinatesRequest) (*pb.GetShipCoordinatesResponse, error) {
	ship, err := s.fleetCache.ShipByName(req.ShipName)
	if err != nil {
		return nil, err
	}
	system, err := s.queries.GetSystemByName(ctx, ship.CurrentLocation.System)
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

	rows, err := s.queries.GetSystemsInRect(ctx, query.GetSystemsInRectParams{
		XMin: int64(rect.Start.X),
		YMin: int64(rect.Start.Y),
		XMax: int64(rect.End.X),
		YMax: int64(rect.End.Y),
	})
	if err != nil {
		return fmt.Errorf("querying systems within rect: %w", err)
	}

	shipMap := s.shipsPerSystem()

	for _, row := range rows {
		system, err := convert.ConvertSystem(&row)
		if err != nil {
			return err
		}
		shipCount := shipMap[system.Id]

		if err = stream.Send(&pb.GetSystemsInRectResponse{
			System:    system,
			ShipCount: int32(shipCount),
		}); err != nil {
			return fmt.Errorf("sending system via gRPC: %w", err)
		}
	}
	return nil
}

func (s *Server) shipsPerSystem() map[string]int {
	m := map[string]int{}

	for _, ship := range s.fleetCache.Ships {
		m[ship.CurrentLocation.System]++
	}
	return m
}
