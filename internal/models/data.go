package models

import "time"

var Chefs = []*Chef{
	{
		ID:        1,
		FirstName: "Gordon",
		LastName:  "Ramsey",
		DOB:       time.Date(1966, 11, 8, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        2,
		FirstName: "Delia",
		LastName:  "Smith",
		DOB:       time.Date(1941, 6, 18, 0, 0, 0, 0, time.UTC),
	},
}

var Meals = []*Meal{
	{
		ID:     1,
		Name:   "Beef Wellington",
		Rating: 5,
		Ingredients: []*Ingredient{
			Ingredients["beef"],
			Ingredients["pastry"],
			// &Ingredient{3, "mushrooms"},
			// &Ingredient{4, "garlic"},
			// &Ingredient{5, "pancetta"},
		},
		Chef: Chefs[0],
	},
	{
		ID:     2,
		Name:   "Apricot Scones",
		Rating: 3,
		Ingredients: []*Ingredient{
			Ingredients["flour"],
			Ingredients["eggs"],
			// "milk",
			// "sugar",
			// "cream of tartar",
			// "baking powder",
			// "dried apricots",
		},
		Chef: Chefs[1],
	},
	{
		ID:     3,
		Name:   "Irish Stew",
		Rating: 4,
		Ingredients: []*Ingredient{
			Ingredients["potatoes"],
			Ingredients["beef"],
			// "carrots",
			// "leek",
			// "stock",
			// "guinness",
		},
		Chef: Chefs[1],
	},
	{
		ID:     4,
		Name:   "Passionfruit Pavlova",
		Rating: 5,
		Ingredients: []*Ingredient{
			Ingredients["egg"],
			Ingredients["sugar"],
			// "vanilla extract",
			// "cornflour",
			// "double cream",
			// "raspberries",
			// "passionfruit",
			// "fresh mint",
			// "icing sugar",
		},
		Chef: Chefs[0],
	},
}

var Ingredients = map[string]*Ingredient{
	"beef":     {1, "beef"},
	"pastry":   {2, "pastry"},
	"flour":    {3, "flour"},
	"egg":      {4, "egg"},
	"potatoes": {5, "potatoes"},
	"sugar":    {6, "sugar"},
}
