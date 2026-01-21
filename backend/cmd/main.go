package main

import "backend/internal/router"

func main() {
	// cfg := config.LoadEnv()
	// connection := database.DbConnection(cfg.Database)
	// fmt.Println(connection)
	// fmt.Println(cfg.Port)

	r := router.AppRouter()
	r.Run()

}
