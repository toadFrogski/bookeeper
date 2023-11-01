package utils

type Response[T any] struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_messsage"`
	Data            T      `json:"data"`
}

func BuildResponse[T any](status ResponseStatus, data T) Response[T] {
	return _BuildResponse[T](status.GetResponseStatus(), status.GetResponseMessage(), data)
}

func _BuildResponse[T any](code int, message string, data T) Response[T] {
	return Response[T]{
		ResponseCode:    code,
		ResponseMessage: message,
		Data:            data,
	}
}
