package dto

import (
	"gg/utils/constants"
)

type Response[T any] struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_messsage"`
	Data            T      `json:"data"`
}

func BuildResponse[T any](status constants.ResponseStatus, data T) Response[T] {
	return Response[T]{
		ResponseCode:    status.GetResponseStatus(),
		ResponseMessage: status.GetResponseMessage(),
		Data:            data,
	}
}
