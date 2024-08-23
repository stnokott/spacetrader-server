package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/stnokott/spacetrader-server/internal/config"
)

func BenchmarkIndex(b *testing.B) {
	log.SetLevel(log.DebugLevel)

	_ = godotenv.Load(".env")
	// load config
	cfg, err := config.Load()
	if err != nil {
		b.Fatal(err)
	}

	dbFile, err := os.CreateTemp("", "space_traders_galaxy*.db")
	if err != nil {
		b.Fatal(err)
	}
	_ = dbFile.Close()
	defer os.Remove(dbFile.Name())

	// create server
	s, err := New(baseURL, cfg.AgentToken, dbFile.Name())
	if err != nil {
		b.Fatal(err)
	}
	defer func() {
		_ = s.Close()
	}()
	b.ResetTimer()
	if err := s.UpdateSystemIndex(true); err != nil {
		b.Fatal(err)
	}
}
