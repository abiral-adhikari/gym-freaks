package handlers

import (
	"encoding/json"
	"gym-freaks-backend/controllers"
	"gym-freaks-backend/models"
	"io"
	"net/http"
)

type FoodHandler struct{}

func (h *FoodHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var food models.Food

	// Read the request body to get token payload
	userDataFromToken, err := controllers.GetTokenPayloadFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// read request to create food
	req, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(req, &food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the user data
	if food.Name == "" || food.Calories <= 0 || food.Unit == "" {
		http.Error(w, "Fields are empty", http.StatusBadRequest)
		return
	}

	controllers.FoodControllers.Create(food)

}
