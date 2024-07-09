// Package main runs a mock server returning static responses for testing.
package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"
	"time"

	pb "github.com/stnokott/spacetrader/internal/proto"
	"github.com/stnokott/spacetrader/tests/mocks"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	port = 55555
)

type MockServer struct {
	pb.UnimplementedSpaceTradersServiceServer
}

func (s *MockServer) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

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

func (s *MockServer) GetCurrentAgent(_ context.Context, _ *emptypb.Empty) (*pb.Agent, error) {
	return &pb.Agent{
		Name:         "STNOKOTT",
		Credits:      math.MaxInt64,
		Headquarters: "GITHUB",
		ShipCount:    42,
	}, nil
}

func (s *MockServer) GetFleet(_ context.Context, _ *emptypb.Empty) (*pb.Fleet, error) {
	return &pb.Fleet{Ships: []*pb.Ship{
		mocks.NewDefaultShip(),
	}}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("TCP listen: %v", err)
		return
	}

	s := &MockServer{}

	srv := grpc.NewServer()
	pb.RegisterSpaceTradersServiceServer(srv, s)
	log.Println("server running on port", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("TCP serve: %v", err)
		return
	}
}
