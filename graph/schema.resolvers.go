package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ypeckstadt/golang-graphql-study/pkg/graph/generated"
	"github.com/ypeckstadt/golang-graphql-study/pkg/graph/model"
)

func (r *myMutationResolver) CreateUser(ctx context.Context, user model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myMutationResolver) UpdateUser(ctx context.Context, uuid string, update model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myMutationResolver) DeleteUser(ctx context.Context, uuid string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myQueryResolver) User(ctx context.Context, uuid string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *myQueryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// MyMutation returns generated.MyMutationResolver implementation.
func (r *Resolver) MyMutation() generated.MyMutationResolver { return &myMutationResolver{r} }

// MyQuery returns generated.MyQueryResolver implementation.
func (r *Resolver) MyQuery() generated.MyQueryResolver { return &myQueryResolver{r} }

type myMutationResolver struct{ *Resolver }
type myQueryResolver struct{ *Resolver }
