package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.53

import (
	"context"

	"github.com/stnokott/spacetrader-server/internal/graph/loaders"
	"github.com/stnokott/spacetrader-server/internal/graph/model"
)

// From is the resolver for the from field.
func (r *jumpgateResolver) From(ctx context.Context, obj *model.Jumpgate) (*model.Waypoint, error) {
	return loaders.GetWaypoint(ctx, obj.FromWaypointID)
}

// To is the resolver for the to field.
func (r *jumpgateResolver) To(ctx context.Context, obj *model.Jumpgate) (*model.Waypoint, error) {
	return loaders.GetWaypoint(ctx, obj.ToWaypointID)
}

// Jumpgate returns JumpgateResolver implementation.
func (r *Resolver) Jumpgate() JumpgateResolver { return &jumpgateResolver{r} }

type jumpgateResolver struct{ *Resolver }
