package middlewares

import (
	"gg/utils/constants"
	"gg/utils/dto"
	"gg/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleAccessMiddleware(accessRoles []constants.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exist := c.Get("user")
		if !exist {
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[any](constants.InternalError, nil))
			c.Abort()
			return
		}
		userRoles := user.(token.Claims).Roles

		for _, accessRole := range accessRoles {
			for _, userRole := range userRoles {
				if string(accessRole) == string(userRole) {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, dto.BuildResponse[any](constants.Unauthorized, nil))
		c.Abort()
	}
}
