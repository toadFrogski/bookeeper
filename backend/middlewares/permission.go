package middlewares

import (
	"bookeeper/domain"
	"bookeeper/utils/constants"
	"bookeeper/utils/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleAccessMiddleware(accessRoles []constants.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *domain.User
		vars, exist := c.Get("user")
		if !exist {
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[any](constants.InternalError, nil))
			c.Abort()
			return
		}
		user = vars.(*domain.User)

		for _, accessRole := range accessRoles {
			for _, userRole := range user.Roles {
				if accessRole == userRole.Name {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, dto.BuildResponse[any](constants.PermissionDenied, nil))
		c.Abort()
	}
}
