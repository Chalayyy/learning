package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/khan/graphql-examples/graph/generated"
	"github.com/khan/graphql-examples/graph/model"
)

func (r *mutationResolver) Increment(ctx context.Context) (int, error) {
	return Increment(), nil
}

func (r *queryResolver) Greeting(ctx context.Context) (string, error) {
	return "Hello, world!", nil
}

func (r *queryResolver) Farewell(ctx context.Context) (string, error) {
	return "See ya!", nil
}

func (r *queryResolver) MyName(ctx context.Context) (string, error) {
	return "Ben", nil
}

func (r *queryResolver) MyAge(ctx context.Context, asOf string) (int, error) {
	return Age(asOf, 1990)
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	index := id - 1
	if index >= len(users) || index < 0 {
		return nil, fmt.Errorf("no user exists with that ID")
	}
	user := users[index]
	return &model.User{
		ID:   id,
		Name: user.name,
	}, nil
}

func (r *userResolver) Age(ctx context.Context, obj *model.User, asOf string) (int, error) {
	return Age(asOf, users[obj.ID-1].birthYear)
}

func (r *userResolver) BestFriend(ctx context.Context, obj *model.User) (*model.User, error) {
	friends := users[obj.ID-1].friendIDs
	if len(friends) == 0 {
		return nil, nil
	}
	bestFriend := users[friends[0]]
	return &model.User{
		ID:   friends[0] + 1,
		Name: bestFriend.name,
	}, nil
}

func (r *userResolver) Friends(ctx context.Context, obj *model.User) ([]model.SocialAnimal, error) {
	user := users[obj.ID-1]

	ret := make([]model.SocialAnimal, 0)
	for _, i := range user.friendIDs {
		ret = append(ret, &model.User{
			ID:   i + 1,
			Name: users[i].name,
		})
	}
	for _, i := range user.dogIDs {
		ret = append(ret, &model.Dog{
			Name: dogs[i].name,
		})
	}
	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
