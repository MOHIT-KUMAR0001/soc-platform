package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"fmt"
)

func main() {
	cfg := config.LoadEnv()
	connection := database.DbConnection(cfg.Database)
    fmt.Println(connection)
	fmt.Println(cfg.Port)
}

