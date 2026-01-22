package router

import (
	"backend/internal/api"
	"backend/internal/database"

	"github.com/gin-gonic/gin"
)

type Apc struct {
	DB *database.Queries
}

func AppRouter() *gin.Engine {
	router := gin.Default()

	g1 := router.Group("/api/v1")

	{
		g1.POST("/ingest", api.App)
	}

	return router

}
