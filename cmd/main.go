// Package main is the gRPC server.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/stnokott/spacetrader-server/internal/config"
)

// define program hooks, mainly used for injecting mocked logic into the main lifecycle
var beforeHooks []func() error
var afterHooks []func() error

var baseURL = "!UNDEFINED"

func main() {
	var err error
	defer func() {
		fmt.Println(err)
		if err != nil {
			os.Exit(1)
		}
	}()

	// run hooks
	for _, h := range beforeHooks {
		if err = h(); err != nil {
			return
		}
	}
	defer func() {
		for _, h := range afterHooks {
			if errInner := h(); errInner != nil {
				fmt.Printf("WARNING: error running hook: %v\n", errInner)
				// we always want to run all after hooks regardless of error so any running background
				// services are always properly closed. This is why we continue on non-nil error here.
			}
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := s.Listen(ctx, 55555, "/graphql"); err != nil {
		logger.Fatal(err)
	}
	// TODO: configure port from env
}
