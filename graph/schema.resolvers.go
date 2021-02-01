package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/daidv2/graphql-go-postgres/database"
	"github.com/daidv2/graphql-go-postgres/graph/generated"
	"github.com/daidv2/graphql-go-postgres/graph/model"
)

func (r *mutationResolver) CreateMember(ctx context.Context, input model.NewMember) (*model.Member, error) {
	member := &model.Member{
		Name: input.Name,
	}
	r.members = append(r.members, member)

	// insert into database
	var lastInsertId int
	err := database.DB.QueryRow("INSERT INTO members(name) VALUES($1) returning id;", input.Name).Scan(&lastInsertId)
	checkErr(err)

	member.ID = fmt.Sprintf("%v", lastInsertId)

	return member, nil
}

func (r *mutationResolver) CreateSkill(ctx context.Context, input model.NewSkill) (*model.Skill, error) {
	skill := &model.Skill{
		Category: input.Category,
		Name:     input.Name,
		Exp:      input.Exp,
	}
	r.skills = append(r.skills, skill)

	// insert into database
	var lastInsertId int
	err := database.DB.QueryRow("INSERT INTO skills(category, name, exp, member_id) VALUES($1, $2, $3, $4) returning id;",
		input.Category,
		input.Name,
		input.Exp,
		input.MemberID,
	).Scan(&lastInsertId)
	checkErr(err)

	skill.ID = fmt.Sprintf("%v", lastInsertId)

	return skill, nil
}

func (r *queryResolver) Members(ctx context.Context, input *model.Params) ([]*model.Member, error) {
	var stmt string
	if input.Ids != nil {
		stmt = "SELECT id, name FROM members WHERE id IN (" + strings.Join(input.Ids, ",") + ")"
	} else {
		stmt = "SELECT id, name FROM members"
	}

	rows, err := database.DB.Query(stmt)
	checkErr(err)
	var members []*model.Member

	for rows.Next() {
		member := &model.Member{}

		err = rows.Scan(&member.ID, &member.Name)
		checkErr(err)
		members = append(members, member)
	}

	return members, nil
}

func (r *queryResolver) Skills(ctx context.Context, input *model.Params) ([]*model.Skill, error) {
	var stmt string
	if input.Ids != nil {
		stmt = "SELECT id, category, name, exp FROM skills WHERE id IN (" + strings.Join(input.Ids, ",") + ")"
	} else {
		stmt = "SELECT id, category, name, exp FROM skills"
	}

	rows, err := database.DB.Query(stmt)
	checkErr(err)
	var skills []*model.Skill

	for rows.Next() {
		skill := &model.Skill{}

		err = rows.Scan(&skill.ID, &skill.Category, &skill.Name, &skill.Exp)
		checkErr(err)
		skills = append(skills, skill)
	}

	return skills, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
