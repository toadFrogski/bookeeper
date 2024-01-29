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

func PanicWithMessage(status constants.ResponseStatus, message string) {
	_PanicException(status.GetResponseStatus(), message)
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		code, _ := strconv.Atoi(strArr[0])
		message := strArr[1]

		if code < 1000 {
			c.JSON(http.StatusInternalServerError, dto.BuildResponse[any](constants.InternalError, nil))
			c.Abort()
			return
		}

		switch code {
		case constants.DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest,
				dto.BuildResponse[any](constants.DataNotFound, nil))
			c.Abort()
		case constants.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized,
				dto.BuildResponse[any](constants.Unauthorized, nil))
			c.Abort()
		case constants.InvalidRequest.GetResponseStatus():
			c.JSON(http.StatusBadRequest,
				dto.BuildResponse[string](constants.InvalidRequest, message))
			c.Abort()
		case constants.UnknownError.GetResponseStatus():
			c.JSON(http.StatusInternalServerError,
				dto.BuildResponse[any](constants.UnknownError, nil))
			c.Abort()
		case constants.InternalError.GetResponseStatus():
			c.JSON(http.StatusInternalServerError,
				dto.BuildResponse[any](constants.InternalError, nil))
			c.Abort()
		default:
			c.JSON(http.StatusBadRequest,
				dto.Response[any]{ResponseCode: code, ResponseMessage: message, Data: nil})
			c.Abort()
		}
	}
}
