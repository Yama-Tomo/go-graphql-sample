package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sample/graph/generated"
	"sample/graph/model"
)

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	name := "Bob"
	return &model.User{
		ID:   id,
		Name: &name,
	}, nil
}

func (r *queryResolver) Pet(ctx context.Context, id *string) (*model.Pet, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
