package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/stnokott/spacetrader-server/internal/config"
)

type cacheWrapper func(ctx context.Context, progressFunc progressFunc) error

func (w cacheWrapper) Create(ctx context.Context, _ *Server, progressFunc progressFunc) error {
	return w(ctx, progressFunc)
}

func TestCacheManager(t *testing.T) {
	shortRunning := func(_ context.Context, progressFunc progressFunc) error {
		progressFunc(1, 0)
		time.Sleep(100 * time.Millisecond)
		progressFunc(1, 1)
		return nil
	}

	longRunning := func(ctx context.Context, progressFunc progressFunc) error {
		i := 0
		target := 10
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				time.Sleep(250 * time.Millisecond)
				i++
				progressFunc(target, i)
				if i == target {
					return nil
				}
			}
		}
	}

	cm := cacheManager{
		caches: map[string]cache{
			"Short": cacheWrapper(shortRunning),
			"Long":  cacheWrapper(longRunning),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := cm.Create(ctx, nil)
	assert.NoError(t, ctx.Err())
	assert.NoError(t, err)
}

func benchIndex(b *testing.B, getIndexFunc func(*Server) func(context.Context) error) {
	log.SetLevel(log.InfoLevel)

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
