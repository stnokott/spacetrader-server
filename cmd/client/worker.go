package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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
	ServerInfo *TypedBinding[*pb.ServerStatusReply]
}

// NewWorker creates a new worker instance.
// A connection is not established until:
// TODO
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

// CheckAppServer pings the app server and returns any error it encounters.
func (w *Worker) CheckAppServer(ctx context.Context) error {
	_, err := w.client.Ping(ctx, nil)
	return err
}

// CheckGameServer queries the current game server status and returns any error it encounters.
//
// It also updates the server status binding on success.
func (w *Worker) CheckGameServer(ctx context.Context) error {
	status, err := w.client.GetServerStatus(ctx, nil)
	if err != nil {
		return err
	}
	w.bindings.ServerInfo.Set(status)
	return nil
}

// Close closes all gRPC connections.
func (w *Worker) Close() error {
	return w.conn.Close()
}
