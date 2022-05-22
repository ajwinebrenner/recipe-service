package resolvers

import "gorm.io/gorm"

type RootResolver struct {
	*ChefResolver
	*MealResolver
	*IngredientResolver
}

func NewRootResolver(db *gorm.DB) *RootResolver {
	return &RootResolver{
		ChefResolver:       NewChefResolver(db),
		MealResolver:       NewMealResolver(db),
		IngredientResolver: NewIngredientResolver(db),
	}
}
