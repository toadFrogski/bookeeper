package middlewares

import (
	"fmt"
	"gg/utils/constants"
	"gg/utils/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO ROLE ASSERTION ON TUPLE
func RoleAccessMiddleware(accessRole constants.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exist := c.Get("user")
		if !exist {
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[any](constants.InternalError, nil))
			c.Abort()
		}
		roles, exist := user.(map[string]any)["roles"].([]any)
		if !exist {
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[any](constants.InternalError, nil))
			c.Abort()
		}
		for _, role := range roles {
			name := role.(string) == string(accessRole)
			fmt.Printf("%x%x", name)
		}
	}
}
