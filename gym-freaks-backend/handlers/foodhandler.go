package handlers

import (
	"encoding/json"
	"fmt"
	"gym-freaks-backend/controllers"
	"gym-freaks-backend/models"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FoodHandler struct{}

// Food Handler to Create Food
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

func Update(w http.ResponseWriter, r *http.Request) {
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

	isCreator := controllers.CheckCreator(&food, userDataFromToken.UserID)
	if !isCreator {
		http.Error(w, "You are donot have authority", http.StatusMethodNotAllowed)
	}

	foodid, err := controllers.FoodControllers.Update(food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message := fmt.Sprintf("Food  with id %v Updated", foodid)
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

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	// Read the request body to get token payload
	userDataFromToken, err := controllers.GetTokenPayloadFromRequest(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	food, err := controllers.FoodControllers.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	isCreator := controllers.CheckCreator(&food, userDataFromToken.UserID)
	if !isCreator {
		http.Error(w, "You are donot have authority", http.StatusMethodNotAllowed)
	}
	foodid, err := controllers.FoodControllers.Delete(food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	message := fmt.Sprintf("Food  with id %v Deleted", foodid)
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

func Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var food models.FoodSearchRequest
	// read request to create food
	req, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(req, &food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	foodlist, err := controllers.FoodControllers.Search(food.FoodName, food.Unit, food.CreatedBy, food.MinCalories, food.MaxCalories)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message := map[string]any{"status": "0", "foods": foodlist}
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

func GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	food, err := controllers.FoodControllers.GetOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message := map[string]any{"status": "0", "food": food}
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
