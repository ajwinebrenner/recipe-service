package models

type Meal struct {
	ID          int
	Name        string
	Rating      int32
	Ingredients []*Ingredient `gorm:"many2many:meal_ingredients;"`
	ChefID      int           // for gorm to referrence foreign key
	Chef        *Chef
}
