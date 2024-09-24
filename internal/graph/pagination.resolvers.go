package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.53

import (
	"github.com/nrfta/go-paging"
)

// PageInfo returns PageInfoResolver implementation.
func (r *Resolver) PageInfo() PageInfoResolver {
	return paging.NewPageInfoResolver()
}
