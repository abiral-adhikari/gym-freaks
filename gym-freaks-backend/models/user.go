package models

import (
	"time"
)

type Role string

const (
	Trainer Role = "trainer"
	Gymer   Role = "gymer"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     int       `json:"phone"`
	Dob       Date      `json:"dob"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}

type LoginData struct {
	Phone    int    `json:"phone"`
	Password string `json:"password"`
}
