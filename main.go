// Package main runs the core program loop.
package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/stnokott/spacetrader/internal/client"
	"github.com/stnokott/spacetrader/internal/config"
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
		log.Fatal(cfg)
	}

	// create API client
	cli := client.New(baseURL, cfg.AgentToken)
	status, err := cli.Status()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(status)
}
