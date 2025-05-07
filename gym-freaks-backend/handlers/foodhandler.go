package handlers

import (
	"encoding/json"
	"fmt"
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
	if userDataFromToken.Role != "trainer" {
		http.Error(w, "Only trainers are allowed to create food", http.StatusForbidden)
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
	food.CreatedBy.ID = userDataFromToken.UserID
	foodid, err := controllers.FoodControllers.Create(food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message := fmt.Sprintf("Food Created with id %v", foodid)
	response, err := json.Marshal(message)
	// Convert the response to JSON
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	// Set the response header and write the response
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}


func 