// Package main is the gRPC server.
package main

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/stnokott/spacetrader-server/internal/config"
)

const (
	environment = "debug" // TODO: configure from env
	baseURL     = "https://api.spacetraders.io/v2/"
)

func initLogger() {
	if environment == "prod" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		DisableTimestamp:       false,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
	log.SetOutput(os.Stdout)
}

func main() {
	initLogger()

	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Error(err)
		return
	}

	// create server
	s, err := New(baseURL, cfg.AgentToken, "./galaxy.db")
	if err != nil {
		log.Error(err)
		return
	}
	defer func() {
		_ = s.Close()
	}()
	// TODO: update in background, return ETA if queried
	if err := s.CreateIndexes(context.Background()); err != nil {
		log.Error(err)
		return
	}

	log.Fatal(s.Listen(55555)) // TODO: configure port from env
	// TODO: graceful shutdown
}
