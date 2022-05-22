package resolvers

import (
	"strconv"

	"github.com/ajwinebrenner/recipe-service/internal/models"
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type Meal struct {
	meal *models.Meal
}

func (m *Meal) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(m.meal.ID))
}

func (m *Meal) Name() string {
	return m.meal.Name
}

func (m *Meal) Rating() int32 {
	return m.meal.Rating
}

func (m *Meal) Ingredients() []*Ingredient {
	var ingredients []*Ingredient
	for _, i := range m.meal.Ingredients {
		ingredients = append(ingredients, &Ingredient{ingredient: i})
	}
	return ingredients
}

func (m *Meal) Chef() *Chef {
	return &Chef{
		chef: m.meal.Chef,
	}
}

type MealResolver struct {
	*MealQuery
}

func NewMealResolver(db *gorm.DB) *MealResolver {
	return &MealResolver{
		NewMealQuery(db),
	}
}
