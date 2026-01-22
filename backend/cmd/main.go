package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/router"
	"backend/internal/schema"
	"database/sql"
	"log"
)

func main() {
	cfg := config.LoadEnv()

	dbUrl := cfg.Database

	if dbUrl == "" {
		log.Fatalf("Could't find database config file")
	}

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatalf("Unable to open database connection")
	}

	r := router.AppRouter()

	r.Run()
}
