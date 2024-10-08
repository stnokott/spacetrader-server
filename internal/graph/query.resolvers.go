package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.53

import (
	"context"
	"errors"

	"github.com/nrfta/go-paging"
	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/stnokott/spacetrader-server/internal/convert"
	"github.com/stnokott/spacetrader-server/internal/db/query"
	"github.com/stnokott/spacetrader-server/internal/graph/model"
)

// Server is the resolver for the server field.
func (r *queryResolver) Server(ctx context.Context) (*model.Server, error) {
	result := new(api.Status)
	if err := r.api.Get(ctx, result, "/"); err != nil {
		return nil, err
	}
	converted := convert.ConvertServerStatus(result)
	return converted, nil
}

// Agent is the resolver for the agent field.
func (r *queryResolver) Agent(ctx context.Context) (*model.Agent, error) {
	result := new(struct {
		// for some reason, SpaceTraders decided it's a good idea to wrap the agent
		// info in a useless "data" field.
		Data *api.Agent `json:"data"`
	})
	if err := r.api.Get(ctx, result, "/my/agent"); err != nil {
		return nil, err
	}

	return convert.ConvertAgent(result.Data), nil
}

// Ships is the resolver for the ships field.
func (r *queryResolver) Ships(_ context.Context) ([]*model.Ship, error) {
	if r.fleetCache.Ships == nil {
		return nil, errors.New("fleet cache has not been initialized")
	}
	return r.fleetCache.Ships, nil
}

// Systems is the resolver for the systems field.
func (r *queryResolver) Systems(ctx context.Context, page *paging.PageArgs) (out *model.SystemConnection, err error) {
	defer func() {
		if err != nil {
			out = &model.SystemConnection{
				PageInfo: paging.NewEmptyPageInfo(),
			}
		}
	}()

	var total int64
	if total, err = r.db.GetSystemCount(ctx); err != nil {
		return
	}

	paginator := paging.NewOffsetPaginator(page, total)

	var rows []query.System
	if rows, err = r.db.GetSystemsOffset(ctx, query.GetSystemsOffsetParams{
		Offset: int64(paginator.Offset),
		Limit:  int64(paginator.Limit),
	}); err != nil {
		return
	}

	edges := make([]*model.SystemEdge, len(rows))
	for i, row := range rows {
		edges[i] = &model.SystemEdge{
			Cursor: paging.EncodeOffsetCursor(paginator.Offset + i + 1),
			Node:   convert.ConvertSystem(row),
		}
	}
	out = &model.SystemConnection{
		PageInfo: &paginator.PageInfo,
		Edges:    edges,
	}
	return
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
