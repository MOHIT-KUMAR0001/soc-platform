package auth

import (
	"backend/internal/database"

	"github.com/gin-gonic/gin"
)

type apiConfig struct {
	db *database.Queries
}

func (db *apiConfig) Auth(ctx *gin.Context) {

}
