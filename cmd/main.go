// Package main is the gRPC server.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/stnokott/spacetrader-server/internal/config"
)

const (
	baseURL = "https://api.spacetraders.io/v2/"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			logger.Error(err)
			os.Exit(1)
		}
	}()

	var pprofFile *os.File
	pprofFile, err = os.Create("cpu.pprof")
	if err != nil {
		err = fmt.Errorf("could not create CPU profile: %w", err)
	}
	defer pprofFile.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(pprofFile); err != nil {
		err = fmt.Errorf("could not start CPU profile: %w", err)
	}
	defer pprof.StopCPUProfile()

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
		if errClose := s.Close(); errClose != nil {
			logger.Errorf("server shutdown failed: %v", errClose)
		}
	}()
	// TODO: update in background, return ETA if queried
	if err = s.CreateCaches(context.Background()); err != nil {
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := s.Listen(ctx, 55555, "/graphql"); err != nil {
		logger.Error(err)
		return
	}
	// TODO: configure port from env
}
