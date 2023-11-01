package utils

type ResponseStatus int

const (
	Success ResponseStatus = iota
	Unauthorized
	DataNotFound
	InvalidRequest
	UnknownError
)

func (r ResponseStatus) GetResponseStatus() int {
	return [...]int{1000, 1001, 1002, 1003, 1004}[r]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"SUCCESS", "UNAUTHORIZED", "DATA_NOT_FOUND", "INVALID_REQUEST", "UNKNOWN_ERROR"}[r]
}
