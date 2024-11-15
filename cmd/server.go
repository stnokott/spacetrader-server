package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/cache"
	"github.com/stnokott/spacetrader-server/internal/graph"
	"github.com/stnokott/spacetrader-server/internal/graph/loaders"
	"github.com/stnokott/spacetrader-server/internal/log"
	"github.com/stnokott/spacetrader-server/internal/worker"

	"github.com/stnokott/spacetrader-server/internal/db/query"
)

var logger = log.ForComponent("server")

// Server performs requests to the SpaceTraders API and offers them via gRPC.
type Server struct {
	api     *api.Client
	db      *sql.DB
	queries *query.Queries
	// TODO: check if api or query can be removed

	systemCache   cache.SystemCache
	jumpgateCache cache.JumpgateCache
	fleetCache    *cache.FleetCache
}

// New creates and returns a new Client instance.
func New(baseURL string, token string, dbFile string) (*Server, error) {
	client := api.NewClient(baseURL, token)

	db, err := newDB(dbFile)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	q, err := query.Prepare(ctx, db)
	if err != nil {
		return nil, err
	}

	return &Server{
		api:     client,
		db:      db,
		queries: q,

		systemCache:   cache.NewSystemCache(client, db, q),
		jumpgateCache: cache.NewJumpgateCache(client, db, q),
		fleetCache:    cache.NewFleetCache(client),
	}, nil
}

// Close terminates all underlying connections.
func (s *Server) Close() error {
	logger.Info("closing database connection")
	return errors.Join(s.queries.Close(), s.db.Close())
}

// Listen starts the GraphQL server on the specified port and path (i.e. /graphql).
//
// This function blocking.
// When the context expires, the server will attempt to shutdown gracefully.
func (s *Server) Listen(ctx context.Context, port int, path string) error {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(
		s.api, s.queries, s.fleetCache,
	)}))
	srv.AddTransport(&transport.Websocket{})

	var srvHandler http.Handler = srv
	// wrap dataloader middleware for injecting dataloaders into request contexts
	srvHandler = loaders.Middleware(s.queries, srvHandler)

	mux := http.NewServeMux()
	mux.Handle(path, srvHandler)
	mux.Handle("/playground", playground.Handler("Playground", path))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	serveErrChan := make(chan error, 0)
	go func() {
		logger.Infof("GraphQL server listening on :%d%s", port, path)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			serveErrChan <- err
			close(serveErrChan)
		}
	}()

	select {
	case <-ctx.Done():
		logger.Info("shutting down GraphQL server")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(ctx)
	case err := <-serveErrChan:
		return err
	}
}

// TODO: write functions for querying API stuff (e.g. getSystem)
// which wrap API calls, but also handle caching.
// So when calling getSystem and the queried system doesn't exist in the cache yet,
// we query the API and write the result to the cache.

var indexTimeout = 1 * time.Hour

// CreateCaches updates or creates all registered indexes.
// It should be called once at the beginning of the program loop.
func (s *Server) CreateCaches(ctxParent context.Context) error {
	logger.Info("creating caches")

	ctx, cancel := context.WithTimeout(ctxParent, indexTimeout)
	defer cancel()

	if err := worker.AddAndWait(ctx, "create-system-cache", func(ctx context.Context, progressChan chan<- float64) error {
		return s.systemCache.Create(ctx, progressChan)
	}, worker.WithMaxLogFrequency(5*time.Second)); err != nil {
		return err
	}
	if err := worker.AddAndWait(ctx, "create-jumpgate-cache", func(ctx context.Context, progressChan chan<- float64) error {
		return s.jumpgateCache.Create(ctx, progressChan)
	}, worker.WithMaxLogFrequency(5*time.Second)); err != nil {
		return err
	}
	if err := worker.AddAndWait(ctx, "create-fleet-cache", func(ctx context.Context, progressChan chan<- float64) error {
		return s.fleetCache.Create(ctx, progressChan)
	}); err != nil {
		return err
	}

	return nil
}
