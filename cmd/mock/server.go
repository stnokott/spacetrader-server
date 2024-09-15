// Package main runs a mock server returning static responses for testing.
package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"time"

	pb "github.com/stnokott/spacetrader-server/internal/proto"
	"github.com/stnokott/spacetrader-server/tests/mocks"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	port = 55555
)

// MockServer returns mocked results for all gRPC services.
type MockServer struct {
	pb.UnimplementedSpacetraderServer

	systems []*pb.System
	ships   []*pb.Ship
}

func newMockServer() *MockServer {
	systems := mocks.GenerateSystems(2000, -10000, 10000)

	systemIdx1 := rand.Intn(len(systems))
	system1 := systems[systemIdx1]
	ship1 := mocks.NewShipInSystem("Enterprise", system1)

	systemIdx2 := rand.Intn(len(systems))
	system2 := systems[systemIdx2]
	ship2 := mocks.NewShipInSystem("Pod Racer", system2)

	return &MockServer{
		systems: systems,
		ships:   []*pb.Ship{ship1, ship2},
	}
}

// Ping is a mock.
func (s *MockServer) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// GetServerStatus is a mock.
func (s *MockServer) GetServerStatus(_ context.Context, _ *emptypb.Empty) (*pb.ServerStatus, error) {
	return &pb.ServerStatus{
		Version:   "v1.2.3",
		LastReset: timestamppb.New(time.Now().Add(-5 * 24 * time.Hour)),
		NextReset: timestamppb.New(time.Now().Add(1 * 24 * time.Hour)),
		GlobalStats: &pb.ServerStatus_GlobalStats{
			Agents:    123,
			Ships:     456,
			Waypoints: 999999,
			Systems:   12345678,
		},
		Announcements: []*pb.ServerStatus_Announcement{
			{
				Title: "Foo expected to return",
				Body:  "Foo is expected to make a fuzzy return!",
			},
			{
				Title: "Bar in the near future",
				Body:  "The infamous Bar will visit our system soon!",
			},
		},
	}, nil
}

// GetCurrentAgent is a mock.
func (s *MockServer) GetCurrentAgent(_ context.Context, _ *emptypb.Empty) (*pb.Agent, error) {
	return &pb.Agent{
		Name:         "STNOKOTT",
		Credits:      math.MaxInt64,
		Headquarters: "GITHUB",
		ShipCount:    42,
	}, nil
}

// GetFleet is a mock.
func (s *MockServer) GetFleet(_ context.Context, _ *emptypb.Empty) (*pb.Fleet, error) {
	return &pb.Fleet{Ships: s.ships}, nil
}

// GetAllSystems is a mock.
func (s *MockServer) GetAllSystems(_ *emptypb.Empty, stream pb.Spacetrader_GetAllSystemsServer) error {
	for _, system := range s.systems {
		if err := stream.Send(&pb.GetAllSystemsResponseItem{
			Name: system.Id,
			Pos: &pb.Vector{
				X: system.X,
				Y: system.Y,
			},
			HasJumpgates: rand.Float32() > 0.5,
			ShipCount:    rand.Int31n(3),
		}); err != nil {
			return fmt.Errorf("sending system: %w", err)
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("TCP listen: %v", err)
		return
	}

	s := newMockServer()

	srv := grpc.NewServer()
	pb.RegisterSpacetraderServer(srv, s)
	log.Println("server running on port", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("TCP serve: %v", err)
		return
	}
}
