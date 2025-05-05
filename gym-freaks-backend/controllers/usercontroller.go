package controllers

import (
	"context"
	"fmt"
	database "gym-freaks-backend/connections"
	"gym-freaks-backend/models"
	"gym-freaks-backend/queries"
	"time"
)

func SignUp(user models.User) (string, error) {
	conn := database.DBConnect()

	// Use the query from the queries package and return the user ID
	var userID int
	var err error

	hash, err := HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	err = conn.QueryRow(context.Background(), queries.CreateUserQuery, user.Username, hash, user.Email, user.Phone, user.Dob.ToTime(), user.Role).Scan(&userID)
	if err != nil {
		return "", fmt.Errorf("error inserting user: %v", err)
	}

	// Update the user model with the generated userID
	user.ID = userID

	// Generate token for the user
	token, err := CreateJWT(user)
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}
	user.Token = token

	// Update the token value in the database
	_, err = conn.Exec(context.Background(), queries.UpdateUserTokenQuery, user.Token, user.ID)
	if err != nil {
		return "", fmt.Errorf("error updating user token: %v", err)
	}

	return user.Token, nil
}

func Login(loginData models.LoginData) (message string, token string, flag bool) {
	conn := database.DBConnect()
	var userfromDB models.User

	// Get row/data from database
	data := conn.QueryRow(context.Background(), queries.LoginQuery, loginData.Phone)
	var dob time.Time
	err := data.Scan(
		&userfromDB.ID, &userfromDB.Username, &userfromDB.Password, &userfromDB.Email, &userfromDB.Phone, &dob, &userfromDB.Role, &userfromDB.CreatedAt, &userfromDB.Token,
	)
	userfromDB.Dob = models.Date(dob) // Convert time.Time to models.Date
	if err != nil {
		return err.Error(), "", false // Return false if user is not found or an error occurs
	}

	// Verify the password
	if !CheckPasswordHash(userfromDB.Password, loginData.Password) {
		return "Password Didnot Match", "", false // Return false if password does not match
	}

	// Generate token for the user
	token, err = CreateJWT(userfromDB)
	if err != nil {
		return "Token Generation Failed", "", false // Return false if token generation fails
	}
	// Update the token value in the database
	_, err = conn.Exec(context.Background(), queries.UpdateUserTokenQuery, token, userfromDB.ID)
	if err != nil {
		return "Token Update Failed", "", false
	}

	return "Successful Login", token, true // Return the token and true if login is successful
}

func Logout(userid int) error {
	conn := database.DBConnect()
	_, err := conn.Exec(context.Background(), queries.LogoutQuery, userid)
	return err
}

func GetUser(userID int) (models.User, error) {
	conn := database.Connect()
	var user models.User
	err := conn.QueryRow(context.Background(), queries.GetUserByIDQuery, userID).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.Dob, &user.Role, &user.CreatedAt)
	if err != nil {
		return models.User{}, fmt.Errorf("error fetching user: %v", err)
	}
	return user, nil
}
