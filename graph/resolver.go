package graph

import (
	"github.com/daidv2/graphql-go-postgres/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	members []*model.Member
	skills []*model.Skill
}
