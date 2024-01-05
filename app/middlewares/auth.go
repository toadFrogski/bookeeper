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
		userID, err := token.ExtractTokenID(c)
		c.Set("user_id", userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.BuildResponse[any](constants.Unauthorized, nil))
			c.Abort()
			return
		}
		c.Next()
	}
}
