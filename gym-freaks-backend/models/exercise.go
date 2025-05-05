package models

import "time"

type Exercise struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type Workout struct {
	ID        int       `json:"id"`
	User      *User     `json:"user"`
	Exercise  *Exercise `json:"exercise"`
	Sets      int       `json:"sets"`
	Reps      int       `json:"reps"`
	Weight    int       `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
