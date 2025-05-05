package handlers

import "net/http"

// // UserHandlerInterface groups all user-related routes
// type UserHandlerInterface interface {
// 	Signup(w http.ResponseWriter, r *http.Request)
// 	Login(w http.ResponseWriter, r *http.Request)
// 	Logout(w http.ResponseWriter, r *http.Request)
// }

// FoodHandlerInterface groups all food-related routes
type FoodHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
