package middlewares

import (
	"bookeeper/database"
	"bookeeper/domain"
	"bookeeper/utils/constants"
	"bookeeper/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *domain.User
		var anonUser *domain.User
		var claims *jwt.MapClaims

		anonUser = &domain.User{
			Roles: []*domain.Role{&domain.Role{Name: constants.Anonymous}},
		}

		claims, err := token.ExtractTokenClaims(c)
		if err != nil {
			c.Set("user", anonUser)
			c.Next()
			return
		}

		userID, err := claims.GetSubject()
		if err != nil {
			c.Set("user", anonUser)
			c.Next()
			return
		}

		if err := database.DB.Preload("Roles").First(&user, userID).Error; err != nil {
			c.Set("user", anonUser)
			c.Next()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
