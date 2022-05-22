package resolvers

import (
	"strconv"

	"github.com/ajwinebrenner/recipe-service/internal/models"
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

type MealQuery struct {
	db *gorm.DB
}

func NewMealQuery(db *gorm.DB) *MealQuery {
	return &MealQuery{db: db}
}

func (m *MealQuery) Meals() []*Meal {
	var meals []*Meal
	for _, v := range models.Meals {
		meals = append(meals, &Meal{meal: v})
	}
	return meals
}

type MealArgs struct {
	ID graphql.ID
}

func (m *MealQuery) Meal(args MealArgs) *Meal {
	for _, v := range models.Meals {
		if graphql.ID(strconv.Itoa(v.ID)) == args.ID {
			return &Meal{meal: v}
		}
	}
	return nil
}
