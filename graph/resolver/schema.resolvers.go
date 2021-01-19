package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"sample/ent"
	"sample/graph/generated"
)

func (r *queryResolver) User(ctx context.Context, id *int) (*ent.User, error) {
	if id == nil {
		return nil, errors.New("required id")
	}

	return r.DB.User.Get(ctx, *id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	return r.DB.User.Query().CollectFields(ctx).All(ctx)
}

func (r *queryResolver) Pet(ctx context.Context, id *int) (*ent.Pet, error) {
	if id == nil {
		return nil, errors.New("required id")
	}

	return r.DB.Pet.Get(ctx, *id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
