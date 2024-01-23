package constants

type ResponseStatus int

const (
	Success ResponseStatus = iota
	Unauthorized
	DataNotFound
	InvalidRequest
	InternalError
	UnknownError
)

func (r ResponseStatus) GetResponseStatus() int {
	return [...]int{1000, 1001, 1002, 1003, 1004, 1005}[r]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"SUCCESS", "UNAUTHORIZED", "DATA_NOT_FOUND", "INVALID_REQUEST", "INTERNAL_ERROR", "UNKNOWN_ERROR"}[r]
}
