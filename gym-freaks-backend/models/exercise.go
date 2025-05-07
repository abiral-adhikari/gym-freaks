package models

import "time"

type Exercise struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	CreatedBy   *User  `json:"createdby"`
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

func (e *Exercise) GetCreatorID() int {
	if e.CreatedBy == nil {
		return 0
	}
	return e.CreatedBy.ID
}

func (w *Workout) GetCreatorID() int {
	if w.User == nil {
		return 0
	}
	return w.User.ID
}
