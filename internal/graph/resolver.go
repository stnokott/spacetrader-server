// Package graph contains resolvers for GraphQL vertices and edges.
package graph

import (
	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/cache"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	"github.com/stnokott/spacetrader-server/internal/log"
)

var logger = log.ForComponent("graph")

//go:generate go run github.com/99designs/gqlgen@v0.17.53 generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the root resolver, containing all required data sources.
type Resolver struct {
	api        *api.Client
	db         *query.Queries
	fleetCache *cache.FleetCache
}

// NewResolver create and return a new Resolver instance.
func NewResolver(api *api.Client, db *query.Queries, fleetCache *cache.FleetCache) *Resolver {
	return &Resolver{
		api:        api,
		db:         db,
		fleetCache: fleetCache,
	}
}
