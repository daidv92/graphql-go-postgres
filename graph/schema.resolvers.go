package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dexp/graph/generated"
	"dexp/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateMember(ctx context.Context, input model.NewMember) (*model.Member, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSkill(ctx context.Context, input *model.NewSkill) (*model.Skill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Members(ctx context.Context) ([]*model.Member, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Skills(ctx context.Context) ([]*model.Skill, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
