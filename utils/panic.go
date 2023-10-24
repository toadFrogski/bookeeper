package utils

import (
	"errors"
	"fmt"
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

func PanicException(response ResponseStatus) {
	_PanicException(response.GetResponseStatus(), response.GetResponseMessage())
}

func PanicHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		code, _ := strconv.Atoi(strArr[0])
		msg := strings.Trim(strArr[1], " ")

		switch code {
		case DataNotFound.GetResponseStatus():
			c.JSON(http.StatusBadRequest, _BuildResponse[any](code, msg, nil))
			c.Abort()
		case Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, _BuildResponse[any](code, msg, nil))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, _BuildResponse[any](code, msg, nil))
			c.Abort()
		}
	}
}
