package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/rock16/recovery/graph/generated"
	"github.com/rock16/recovery/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// CreatePoint is the resolver for the createPoint field.
func (r *mutationResolver) CreatePoint(ctx context.Context, input *model.NewPoint) (string, error) {
	panic(fmt.Errorf("not implemented: CreatePoint - createPoint"))
}

// Points is the resolver for the points field.
func (r *queryResolver) Points(ctx context.Context) ([]*model.Point, error) {
	var points []*model.Point
	dummyPoint := model.Point{
		Point: "30",
	}
	points = append(points, &dummyPoint)
	return points, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
