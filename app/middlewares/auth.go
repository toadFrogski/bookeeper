package middlewares

import (
	"gg/utils/constants"
	"gg/utils/dto"
	"gg/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.BuildResponse[any](constants.Unauthorized, nil))
			c.Abort()
			return
		}
		c.Next()
	}
}
