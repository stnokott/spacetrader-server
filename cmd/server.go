package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stnokott/spacetrader/internal/api"
	"github.com/stnokott/spacetrader/internal/convert"
	"google.golang.org/grpc"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/stnokott/spacetrader/internal/proto"
)

// Server performs requests to the SpaceTraders API and offers them via gRPC.
type Server struct {
	api *resty.Client
	db  *sql.DB

	pb.UnimplementedSpacetraderServer
}

// New creates and returns a new Client instance.
func New(baseURL string, token string, dbFile string) (*Server, error) {
	r := resty.New()
	configureRestyClient(r, baseURL, token)

	ctxDB, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	db, err := newDB(ctxDB, dbFile)
	if err != nil {
		return nil, err
	}

	return &Server{
		api: r,
		db:  db,
	}, nil
}

// Close terminates all underlying connections.
func (s *Server) Close() error {
	return s.db.Close()
}

// Listen starts the gRPC server on the specified port.
//
// It is blocking.
func (s *Server) Listen(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("TCP listen: %w", err)
	}
	srv := grpc.NewServer()
	pb.RegisterSpacetraderServer(srv, s)
	log.WithField("port", port).Infof("gRPC server listening")
	if err := srv.Serve(lis); err != nil {
		return fmt.Errorf("TCP serve: %w", err)
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
	if err := s.get(ctx, result, "/", 200); err != nil {
		return nil, err
	}

	return convert.ConvertStatus(result), nil
}

// GetCurrentAgent returns information about the agent identified by the current token.
func (s *Server) GetCurrentAgent(ctx context.Context, _ *emptypb.Empty) (*pb.Agent, error) {
	result := new(struct {
		// for some reason, SpaceTraders decided it's a good idea to wrap the agent
		// info in a useless "data" field.
		Data *api.Agent `json:"data"`
	})
	if err := s.get(ctx, result, "/my/agent", 200); err != nil {
		return nil, err
	}

	return convert.ConvertAgent(result.Data)
}

// GetFleet returns the complete list of ships in the agent's posession.
func (s *Server) GetFleet(ctx context.Context, _ *emptypb.Empty) (*pb.Fleet, error) {
	// we need the total ship count to enable allocation of the correct slice size
	agent, err := s.GetCurrentAgent(ctx, nil)
	if err != nil {
		return nil, err
	}
	out := make([]*pb.Ship, agent.ShipCount)

	// TODO: refactor to use pagination wrapper for all paginated endpoints

	n := 0
	page := 1
	perPage := 10
	for n < len(out) {
		ships, err := s.getFleetPaginated(ctx, page, perPage)
		if err != nil {
			return nil, err
		}

		for _, ship := range ships {
			shipConverted, err := convert.ConvertShip(ship)
			if err != nil {
				return nil, err
			}
			out[n] = shipConverted
			n++
		}
		page++
	}
	return &pb.Fleet{Ships: out}, nil
}

func (s *Server) getFleetPaginated(ctx context.Context, page int, limit int) ([]*api.Ship, error) {
	url := fmt.Sprintf("/my/ships?page=%d&limit=%d", page, limit)

	result := new(struct {
		Data []*api.Ship `json:"data"`
	})
	if err := s.get(ctx, result, url, 200); err != nil {
		return nil, err
	}
	return result.Data, nil
}
