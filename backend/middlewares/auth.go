package middlewares

import (
	"gg/utils/constants"
	"gg/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := token.ExtractTokenClaims(c)
		if user == nil || err != nil {
			user = &token.Claims{
				UserID: 0,
				Roles:  []constants.Role{constants.Anonymous},
			}
		}
		c.Set("user", *user)
		c.Next()
	}
}
