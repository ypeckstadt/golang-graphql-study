package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/ypeckstadt/golang-graphql-study/pkg/graph/generated"
	"github.com/ypeckstadt/golang-graphql-study/pkg/graph/model"
	"log"

	"github.com/ypeckstadt/golang-graphql-study/pkg/server/user"
)

func (r *myMutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	newUser, err := r.User.Create(user.Name)
	if err != nil {
		return nil, err
	}

	log.Printf("User saved with identifier: %d", newUser.ID)
	return &model.User{
		UUID: newUser.UUID,
		Name: newUser.Name,
	}, nil
}

func (r *myMutationResolver) UpdateUser(ctx context.Context, uuid string, update model.UserInput) (*model.User, error) {
	// Update user
	user, err := r.User.Update(uuid, user.UpdateRequest{Name: update.Name})
	if err != nil {
		return nil, err
	}

	log.Printf("User with identifier: %d updated", user.ID)

	// Convert and return via Graph QL
	return &model.User{
		UUID: user.UUID,
		Name: user.Name,
	}, nil
}

func (r *myMutationResolver) DeleteUser(ctx context.Context, uuid string) (bool, error) {
	err := r.User.Delete(uuid)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *myQueryResolver) User(ctx context.Context, uuid string) (*model.User, error) {
	user, err := r.Resolver.User.Get(uuid)
	if err != nil {
		return nil, err
	}
	return &model.User{
		UUID: user.UUID,
		Name: user.Name,
	}, nil
}

func (r *myQueryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	savedUsers, err := r.Resolver.User.List()
	if err != nil {
		return nil, err
	}
	for i, savedUser := range savedUsers {
		var user model.User
		savedUser = savedUsers[i]
		user.UUID = savedUser.UUID
		user.Name = savedUser.Name
		users = append(users, &user)
	}
	return users, nil
}

// MyMutation returns generated1.MyMutationResolver implementation.
func (r *Resolver) MyMutation() generated.MyMutationResolver { return &myMutationResolver{r} }

// MyQuery returns generated1.MyQueryResolver implementation.
func (r *Resolver) MyQuery() generated.MyQueryResolver { return &myQueryResolver{r} }

type myMutationResolver struct{ *Resolver }
type myQueryResolver struct{ *Resolver }
