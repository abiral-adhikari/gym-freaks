package handlers

import (
	"encoding/json"
	"fmt"

	// "go/token"
	controllers "gym-freaks-backend/controllers"
	"gym-freaks-backend/middleware"
	"gym-freaks-backend/models"
	"io"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var user models.User
	// var requestbody map[string]any
	req, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(req, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Println(user)
	// Validate the user data
	if user.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}
	if user.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}
	if user.Phone == 0 {
		http.Error(w, "Phone number is required", http.StatusBadRequest)
		return
	}
	if user.Dob.ToTime().IsZero() {
		http.Error(w, "Date of birth is required", http.StatusBadRequest)
		return
	}
	if !(user.Role == "trainer" || user.Role == "user") {
		http.Error(w, "Role must be either trainer or user", http.StatusBadRequest)
		return
	}

	// Insert to database
	token, err := controllers.SignUp(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := map[string]any{"message": "User Created Successfully", "token": token, "created at ": user.CreatedAt}
	response, err := json.Marshal(payload)
	// Convert the response to JSON
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	// Set the response header and write the response
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}





func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var data map[string]any
	var loginData models.LoginData

	// Read the request body
	request, _ := io.ReadAll(r.Body)
	json.Unmarshal(request, &data)
	fmt.Printf("%T", data["phone"])
	// Validate the request body
	phone, ok := data["phone"].(float64)
	if !ok {
		http.Error(w, "Invalid Payload in field Phone", http.StatusBadRequest)
		return
	}
	password, ok := data["password"].(string)
	if !ok {
		http.Error(w, "Invalid Payload in field Password", http.StatusBadRequest)
		return
	}
	loginData.Phone = int(phone)
	loginData.Password = password

	// Validate the user data
	message, token, flag := controllers.Login(loginData)
	payload := map[string]any{"message": message, "token": token, "flag": flag}
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

}







func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	// Read the request body
	token, err := middleware.GetTokenFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userDataFromToken, err := controllers.VerifyJWT(token)
	// fmt.Printf("%v %v", userDataFromToken.UserID, userDataFromToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userid := userDataFromToken.UserID
	username := userDataFromToken.Username
	// role, ok := userDataFromToken["role"].(string)

	// Update the token value in the database
	err = controllers.Logout(userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload := map[string]any{"message": fmt.Sprintf(`%v Logged Out Successfully `, username), "logoutflag": true}
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
