package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/stnokott/spacetrader/cmd/client/widgets"
	pb "github.com/stnokott/spacetrader/internal/proto"
)

// Worker communicates with the app server in the background.
// The queried data is provided via bindings.
type Worker struct {
	conn   *grpc.ClientConn
	client pb.SpaceTradersServiceClient

	bindings WorkerBindings
}

// WorkerBindings contains all bindings required for a Worker instance.
type WorkerBindings struct {
	Server *widgets.TypedBinding[*pb.ServerStatus]
	Agent  *widgets.TypedBinding[*pb.Agent]
	Fleet  *widgets.TypedBinding[*pb.Fleet]
}

// NewWorker creates a new worker instance.
// A connection is not established until any of the query methods are called.
func NewWorker(addr string, bindings WorkerBindings) *Worker {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// grpc.NewClient only returns errors when the client *setup* is incorrect.
		// It doesn't actually attempt to connect.
		// Thus, we can safely panic here since any error indicates an issue with the code that needs fixing.
		log.Fatalf("gRPC client setup: %v", err)
	}

	client := pb.NewSpaceTradersServiceClient(conn)

	return &Worker{
		conn:   conn,
		client: client,

		bindings: bindings,
	}
}

// Close closes all gRPC connections.
func (w *Worker) Close() error {
	return w.conn.Close()
}

// CheckAppServer pings the app server and returns any error it encounters.
func (w *Worker) CheckAppServer(ctx context.Context) error {
	_, err := w.client.Ping(ctx, nil)
	return err
}

// UpdateServerInfo queries the current game server status and returns any error it encounters.
//
// It also updates the server status binding on success.
func (w *Worker) UpdateServerInfo(ctx context.Context) error {
	status, err := w.client.GetServerStatus(ctx, nil)
	if err != nil {
		return fmt.Errorf("get server status: %w", err)
	}
	w.bindings.Server.Set(status)
	return nil
}

// UpdateCurrentAgent updates the information about the current agent.
func (w *Worker) UpdateCurrentAgent(ctx context.Context) error {
	agent, err := w.client.GetCurrentAgent(ctx, nil)
	if err != nil {
		return fmt.Errorf("get current agent: %w", err)
	}
	w.bindings.Agent.Set(agent)
	return nil
}

// UpdateFleet updates the current agent's ships.
func (w *Worker) UpdateFleet(ctx context.Context) error {
	fleet, err := w.client.GetFleet(ctx, nil)
	if err != nil {
		return fmt.Errorf("get fleet: %w", err)
	}
	w.bindings.Fleet.Set(fleet)
	return nil
}
