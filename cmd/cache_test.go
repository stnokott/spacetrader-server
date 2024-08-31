package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/stnokott/spacetrader-server/internal/config"
)

func benchIndex(b *testing.B, getIndexFunc func(*Server) func(context.Context) error) {
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
	start := time.Now()
	b.ResetTimer()
	if err := getIndexFunc(s)(context.Background()); err != nil {
		b.Fatal(err)
	}
	b.StopTimer()
	b.Logf("duration: %v", time.Now().Sub(start))
}

func BenchmarkIndex(b *testing.B) {
	benchIndex(b, func(s *Server) func(context.Context) error {
		return s.CreateCaches
	})
}
