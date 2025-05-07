package models

import "time"

type Food struct {
	FoodID    int    `json:"id"`
	Name      string `json:"name"`
	Calories  int    `json:"calories"`
	Unit      string `json:"unit"`
	CreatedBy *User  `json:"createdby"`
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

func (f *Food) GetCreatorID() int {
	if f.CreatedBy == nil {
		return 0
	}
	return f.CreatedBy.ID
}

func (m *Meal) GetCreatorID() int {
	if m.User == nil {
		return 0
	}
	return m.User.ID
}
