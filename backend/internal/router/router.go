package router

import (
	"backend/internal/api"

	"github.com/gin-gonic/gin"
)

func AppRouter() *gin.Engine {
	router := gin.Default()

	g1 := router.Group("/api/v1")

	{
		g1.POST("/ingest", api.App)
	}

	return router

}
