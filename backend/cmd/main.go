package main

import (
	"backend/internal/config"
	"backend/internal/router"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func dbConnection(url string) (*sql.DB, error) {
	// 1. Validate the input
	if url == "" {
		return nil, fmt.Errorf("database connection string is empty")
	}

	// 2. Initialize the connection pool
	// Note: sql.Open only validates the format of the connection string
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("could not parse database configuration: %w", err)
	}

	// 3. Verify the connection
	// This actually sends a request to the database to see if it's alive
	if err := db.Ping(); err != nil {
		db.Close() // Clean up if the ping fails
		return nil, fmt.Errorf("database unreachable: %w", err)
	}

	// Success! We return the connection object for use in the app
	log.Println("Successfully established database connection pool")
	return db, nil
}

func main() {
	// Loads the config file
	cfg := config.LoadEnv()

	// calling database connection function
	db, err := dbConnection(cfg.Database)
	if err != nil {
		log.Fatalf("Database setup failed: %v", err)
	}

	// Ensure the database closes when the program exits
	defer db.Close()

	// setting up routers for api
	r := router.AppRouter()
	r.Run()
}
