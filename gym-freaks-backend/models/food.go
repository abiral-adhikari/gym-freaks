package models

import "time"

type Food struct {
	FoodID    int    `json:"id"`
	Name      string `json:"name"`
	Calories  int    `json:"calories"`
	Unit      string `json:"unit"`
	// CreatedBy *User  `json:"createdby"`
}

type Meal struct {
	MealID   int       `json:"id"`
	User     *User     `json:"user"`
	Food     *Food     `json:"food"`
	Quantity int       `json:"quantity"`
	Time     time.Time `json:"time"`
	MealType string    `json:"mealType"`
	Notes    string    `json:"notes"`
}

