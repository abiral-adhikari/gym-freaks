package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"gym-freaks-backend/queries"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var db *pgx.Conn

func Connect() *pgx.Conn {
	godotenv.Load()
	fmt.Println("Starting connection to database...")
	fmt.Println(os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal("Error in establishing connection to database: ", err)
	}

	fmt.Println("Connection to database established successfully")
	_, err = conn.Exec(context.Background(), queries.CreateUserTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table created successfully")
	db = conn
	return conn
}

func DBConnect() *pgx.Conn {
	if db == nil {
		log.Fatal("Database connection is not established")
	}
	return db
}

func Close() {
	fmt.Println("Starting disconnection from database...")
	if db == nil {
		fmt.Println("No connection to close")
		return
	} else {
		db.Close(context.Background())
		fmt.Println("Connection to database closed successfully")
	}
}
