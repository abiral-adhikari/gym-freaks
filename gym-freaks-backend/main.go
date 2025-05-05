package main

import (
	"fmt"
	database "gym-freaks-backend/connections"
	router "gym-freaks-backend/routes"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server is Booting....")
	// Connect to the database
	database.Connect()
	defer database.Close()
	router := router.Router()
	fmt.Println("Server running at http://localhost:8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
