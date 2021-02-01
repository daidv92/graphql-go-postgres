package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/daiv2/graphql-go-postgres/graph/generated"
	"github.com/daiv2/graphql-go-postgres/graph/model"
	database "github.com/daiv2/graphql-go-postgres/database"
)

func (r *mutationResolver) CreateMember(ctx context.Context, input model.NewMember) (*model.Member, error) {
	m := &model.Member{
		Name: input.Name,
	}

	// insert into database
	result, err := database.Exec("INSERT INTO `members` (name) VALUES(?)", m.Name)

	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	m.ID = int(lastId)

	return m, nill
}

func (r *mutationResolver) CreateSkill(ctx context.Context, input model.NewSkill) (*model.Skill, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Members(ctx context.Context) ([]*model.Member, error) {
	var res []*model.Member

	user1 := &model.Member{
		ID:   1,
		Name: "DaiDV",
	}

	user2 := &model.Member{
		ID:   2,
		Name: "DinhLV",
	}

	res = append(res, user1)
	res = append(res, user2)
	return res, nil
}

func (r *queryResolver) Skills(ctx context.Context) ([]*model.Skill, error) {
	var res []*model.Skill

	skill1 := &model.Skill{
		Category: "IT",
		Name:     "PHP",
		Exp:      "7",
	}

	skill2 := &model.Skill{
		Category: "IT",
		Name:     "JAVA",
		Exp:      "10",
	}

	skill3 := &model.Skill{
		Category: "Design",
		Name:     "DD",
		Exp:      "2",
	}

	res = append(res, skill1)
	res = append(res, skill2)
	res = append(res, skill3)
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
