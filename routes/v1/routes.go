package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetV1Routes(r *gin.Engine, db *gorm.DB) *gin.Engine {
	v1 := r.Group("/v1")
	GetBooksRoutes(v1, db)
	return r
}
