package routes

import (
	"github.com/gin-gonic/gin"
)

func GetV1Routes(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1")
	GetBooksRoutes(v1)
	return r
}
