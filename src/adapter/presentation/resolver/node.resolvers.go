package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.48

import (
	"context"
	"fmt"

	"github.com/yuorei/bukubukubooking-back/graph/generated"
	"github.com/yuorei/bukubukubooking-back/src/domain"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (domain.Node, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]domain.Node, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
