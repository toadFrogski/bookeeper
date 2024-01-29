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
		var user token.Claims
		claims, exist := c.Get("user")
		if !exist {
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[any](constants.InternalError, nil))
			c.Abort()
			return
		}
		user = claims.(token.Claims)
		userRoles := user.Roles

		for _, accessRole := range accessRoles {
			for _, userRole := range userRoles {
				if accessRole == userRole {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, dto.BuildResponse[any](constants.Unauthorized, nil))
		c.Abort()
	}
}
