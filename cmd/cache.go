package main

import (
	"context"
	"fmt"
	"math"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stnokott/spacetrader/internal/api"
)

var buildSystemIndexTimeout = 5 * time.Minute

// UpdateSystemIndex queries all systems from the API and writes them to the DB.
// This index can be used later to query systems quickly without relying on the API.
// This approach is valid since systems are expected to be static.
//
// This function is blocking.
func (s *Server) UpdateSystemIndex(force bool) error {
	log.WithField("timeout", buildSystemIndexTimeout).Info("building system index")

	ctx, cancel := context.WithTimeout(context.Background(), buildSystemIndexTimeout)
	defer cancel()

	if !force {
		hasIndex, err := s.hasSystems(ctx)
		if err != nil {
			return fmt.Errorf("checking for system index: %w", err)
		}
		if hasIndex {
			log.Info("system index exists, skipping refresh")
			return nil
		}
	} else {
		log.Info("forcing system index refresh")
	}

	// reads (DB Insert) expected to be faster than writes (API calls)
	// => unbuffered channel
	systemChan := make(chan *api.System)
	errChan := make(chan error, 1)

	go s.getSystems(ctx, systemChan, errChan)
	if err := s.replaceSystems(ctx, systemChan); err != nil {
		return err
	}
	if apiErr, ok := <-errChan; ok {
		return apiErr
	}
	return nil
}

func (s *Server) getSystems(ctx context.Context, systemChan chan<- *api.System, errChan chan<- error) {
	log.Debug("querying systems from API")
	defer close(systemChan)
	defer close(errChan)

	total := math.MaxInt
	n := 0
	for page := 1; n < total; page++ {
		url := fmt.Sprintf("/systems?page=%d&limit=20", page)
		result := new(struct {
			Data []*api.System
			Meta *api.Meta
		})
		if err := s.get(ctx, result, url, 200); err != nil {
			log.Error(err)
			errChan <- err
			return
		}
		total = result.Meta.Total
		n += len(result.Data)
		for _, system := range result.Data {
			systemChan <- system
		}
		log.Infof("querying systems %04d/%04d", n, total)
	}
}
