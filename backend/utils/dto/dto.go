package dto

import (
	"bookeeper/utils/constants"
)

type (
	Response[T any] struct {
		ResponseCode    int    `json:"response_code"`
		ResponseMessage string `json:"response_message"`
		Data            T      `json:"data"`
	} // @name Response

	AnyResponse = Response[any] // @name AnyResponse
)

func BuildResponse[T any](status constants.ResponseStatus, data T) Response[T] {
	return Response[T]{
		ResponseCode:    status.GetResponseStatus(),
		ResponseMessage: status.GetResponseMessage(),
		Data:            data,
	}
}
