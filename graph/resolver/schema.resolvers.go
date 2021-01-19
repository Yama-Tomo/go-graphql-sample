package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"sample/ent/user"
	"sample/graph/generated"
	"sample/graph/model"
	"strconv"
)

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	user := &model.User{ID: nil, Name: nil}
	if id == nil {
		return user, nil
	}

	intId, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	row, err := r.DB.User.Get(ctx, intId)
	if err != nil {
		return nil, err
	}

	strId := strconv.Itoa(row.ID)
	user.ID = &strId
	user.Name = &row.Name

	return user, nil
}

func (r *queryResolver) Pet(ctx context.Context, id *string) (*model.Pet, error) {
	pet := &model.Pet{ID: nil, Name: nil}
	if id == nil {
		return pet, nil
	}

	intId, err := strconv.Atoi(*id)
	if err != nil {
		return nil, err
	}

	row, err := r.DB.Pet.Get(ctx, intId)
	if err != nil {
		return nil, err
	}

	strId := strconv.Itoa(row.ID)
	pet.ID = &strId
	pet.Name = &row.Name

	return pet, nil
}

func (r *userResolver) Pets(ctx context.Context, obj *model.User) ([]*model.Pet, error) {
	userId, err := strconv.Atoi(*obj.ID)
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.User.Query().Where(user.ID(userId)).QueryPets().All(ctx)
	if err != nil {
		return nil, err
	}

	var pets []*model.Pet
	for _, pet := range rows {
		strId := strconv.Itoa(pet.ID)
		pets = append(pets, &model.Pet{ID: &strId, Name: &pet.Name})
	}

	return pets, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
