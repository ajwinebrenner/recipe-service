package resolvers

import (
	"strconv"

	"github.com/ajwinebrenner/recipe-service/internal/models"
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type IngredientQuery struct {
	db *gorm.DB
}

func NewIngredientQuery(db *gorm.DB) *IngredientQuery {
	return &IngredientQuery{db: db}
}

func (c *IngredientQuery) Ingredients() []*Ingredient {
	var ingredients []*Ingredient
	for _, v := range models.Ingredients {
		ingredients = append(ingredients, &Ingredient{ingredient: v})
	}
	return ingredients
}

type IngredientArgs struct {
	ID graphql.ID
}

func (c *IngredientQuery) Ingredient(args IngredientArgs) *Ingredient {
	for _, v := range models.Ingredients {
		if graphql.ID(strconv.Itoa(v.ID)) == args.ID {
			return &Ingredient{ingredient: v}
		}
	}
	return nil
}
