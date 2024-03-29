package middlewares

import (
	"bookeeper/utils/panic"

	"github.com/gin-gonic/gin"
)

func PanicHandleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer panic.PanicHandler(c)
		c.Next()
	}
}
