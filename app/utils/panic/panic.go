package panic

import (
	"errors"
	"fmt"
	"gg/utils/constants"
	"gg/utils/dto"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func _PanicException(code int, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%d: %w", code, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(response constants.ResponseStatus) {
	_PanicException(response.GetResponseStatus(), response.GetResponseMessage())
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		status, _ := strconv.Atoi(strArr[0])

		switch status {
		case constants.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, dto.BuildResponse[any](constants.DataNotFound, nil))
			c.Abort()
		case constants.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, dto.BuildResponse[any](constants.Unauthorized, nil))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[any](constants.UnknownError, nil))
			c.Abort()
		}
	}
}
