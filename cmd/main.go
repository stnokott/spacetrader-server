// Package main is the gRPC server.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/stnokott/spacetrader-server/internal/config"
)

const (
	baseURL = "https://api.spacetraders.io/v2/"
)

func main() {
	var err error
	defer func() {
		fmt.Println(err)
		if err != nil {
			os.Exit(1)
		}
	}()

	// load config
	var cfg *config.Config
	cfg, err = config.Load()
	if err != nil {
		config.PrintUsage(os.Stderr)
		return
	}

	// create server
	var s *Server
	s, err = New(baseURL, cfg.AgentToken, "./galaxy.db")
	if err != nil {
		return
	}
	defer func() {
		_ = s.Close()
	}()
	// TODO: update in background, return ETA if queried
	if err = s.CreateCaches(context.Background()); err != nil {
		return
	}

	logger.Fatal(s.Listen(55555)) // TODO: configure port from env
	// TODO: graceful shutdown
}
