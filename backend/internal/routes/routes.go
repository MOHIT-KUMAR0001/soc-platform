package routes

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	{
		auth := api.Group("/auth")
		{
			auth.GET("/")
		}
	}

	return r
}
