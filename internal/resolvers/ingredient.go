package resolvers

import (
	"strconv"

	"github.com/ajwinebrenner/recipe-service/internal/models"
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type Ingredient struct {
	ingredient *models.Ingredient
}

func (i *Ingredient) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(i.ingredient.ID))
}

func (i *Ingredient) Name() string {
	return i.ingredient.Name
}

type IngredientResolver struct {
	*IngredientQuery
}

func NewIngredientResolver(db *gorm.DB) *IngredientResolver {
	return &IngredientResolver{
		NewIngredientQuery(db),
	}
}
