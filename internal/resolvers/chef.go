package resolvers

import (
	"strconv"

	"github.com/ajwinebrenner/recipe-service/internal/models"
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type Chef struct {
	chef *models.Chef
}

func (c *Chef) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(c.chef.ID))
}

func (c *Chef) FirstName() string {
	return c.chef.FirstName
}

func (c *Chef) LastName() string {
	return c.chef.LastName
}

func (c *Chef) DOB() graphql.Time {
	return graphql.Time{Time: c.chef.DOB}
}

type ChefResolver struct {
	*ChefQuery
}

func NewChefResolver(db *gorm.DB) *ChefResolver {
	return &ChefResolver{
		NewChefQuery(db),
	}
}
