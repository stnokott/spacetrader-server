// Package loaders implements dataloaders for avoiding redundant data queries.
package loaders

import (
	"context"
	"net/http"
	"time"

	"github.com/stnokott/spacetrader-server/internal/convert"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	"github.com/stnokott/spacetrader-server/internal/graph/model"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
	loaderWait = 30 * time.Millisecond // timespan for which requests will be batched
)

// systemLoader reads Systems from DB.
type systemLoader struct {
	db *query.Queries
}

func (s *systemLoader) getSystems(ctx context.Context, systemIDs []string) ([]*model.System, []error) {
	systems, err := s.db.GetSystemsByName(ctx, systemIDs)
	if err != nil {
		return nil, []error{err}
	}

	return convert.ConvertSystems(systems), nil
}

// waypointLoader reads Waypoints from DB.
type waypointLoader struct {
	db *query.Queries
}

func (s *waypointLoader) getWaypoints(ctx context.Context, waypointIDs []string) ([]*model.Waypoint, []error) {
	waypoints, err := s.db.GetWaypointsByName(ctx, waypointIDs)
	if err != nil {
		return nil, []error{err}
	}

	return convert.ConvertWaypoints(waypoints), nil
}

type jumpgateLoader struct {
	db *query.Queries
}

func (j *jumpgateLoader) getJumpgates(ctx context.Context, systemIDs []string) ([][]*model.Jumpgate, []error) {
	out := make([][]*model.Jumpgate, len(systemIDs))
	errs := make([]error, len(systemIDs))

	for i, systemID := range systemIDs {
		if jumpgates, err := j.db.GetJumpgatesInSystem(ctx, systemID); err != nil {
			errs[i] = err
		} else {
			out[i] = convert.ConvertJumpgates(jumpgates)
		}
	}

	return out, errs
}

// Loaders wrap data loaders to inject via middleware
type Loaders struct {
	SystemLoader    *dataloadgen.Loader[string, *model.System]
	WaypointsLoader *dataloadgen.Loader[string, *model.Waypoint]
	JumpgatesLoader *dataloadgen.Loader[string, []*model.Jumpgate]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(db *query.Queries) *Loaders {
	// define the data loader
	sr := &systemLoader{db: db}
	wr := &waypointLoader{db: db}
	jr := &jumpgateLoader{db: db}
	return &Loaders{
		SystemLoader:    dataloadgen.NewLoader(sr.getSystems, dataloadgen.WithWait(loaderWait)),
		WaypointsLoader: dataloadgen.NewLoader(wr.getWaypoints, dataloadgen.WithWait(loaderWait)),
		JumpgatesLoader: dataloadgen.NewLoader(jr.getJumpgates, dataloadgen.WithWait(loaderWait)),
	}
}

// Middleware injects data loaders into the context
func Middleware(db *query.Queries, next http.Handler) http.Handler {
	loader := NewLoaders(db)
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetSystem returns single system by id efficiently
func GetSystem(ctx context.Context, systemID string) (*model.System, error) {
	loaders := For(ctx)
	return loaders.SystemLoader.Load(ctx, systemID)
}

// GetWaypoint returns single waypoint by id efficiently
func GetWaypoint(ctx context.Context, waypointID string) (*model.Waypoint, error) {
	loaders := For(ctx)
	return loaders.WaypointsLoader.Load(ctx, waypointID)
}

// GetJumpgates returns the jumpgates in a system efficiently
func GetJumpgates(ctx context.Context, systemID string) ([]*model.Jumpgate, error) {
	loaders := For(ctx)
	return loaders.JumpgatesLoader.Load(ctx, systemID)
}
