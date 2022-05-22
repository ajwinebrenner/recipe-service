package resolvers

import (
	"strconv"

	"github.com/ajwinebrenner/recipe-service/internal/models"
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type ChefQuery struct {
	db *gorm.DB
}

func NewChefQuery(db *gorm.DB) *ChefQuery {
	return &ChefQuery{db: db}
}

func (c *ChefQuery) Chefs() []*Chef {
	var chefs []*Chef
	for _, v := range models.Chefs {
		chefs = append(chefs, &Chef{chef: v})
	}
	return chefs
}

type ChefArgs struct {
	ID graphql.ID
}

func (c *ChefQuery) Chef(args ChefArgs) *Chef {
	for _, v := range models.Chefs {
		if graphql.ID(strconv.Itoa(v.ID)) == args.ID {
			return &Chef{chef: v}
		}
	}
	return nil
}
