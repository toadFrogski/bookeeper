package constants

type ResponseStatus int

const (
	// Common statuses
	Success ResponseStatus = iota
	Unauthorized
	DataNotFound
	InvalidRequest
	InternalError
	UnknownError
	PermissionDenied

	// User statuses
	RegisteredEmail
	UserNotFound
	IncorrectPassword
	RegisteredUsername

	// API statuses
	ValidationError
)

func (r ResponseStatus) GetResponseStatus() int {
	return [...]int{
		// Common status codes
		1000, 1001, 1002, 1003, 1004, 1005, 1006,

		// User status codes
		2001, 2002, 2003, 2004,

		// API status messages,
		3001,
	}[r]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{
		// Common status messages
		"SUCCESS", "UNAUTHORIZED", "DATA_NOT_FOUND", "INVALID_REQUEST",
		"INTERNAL_ERROR", "UNKNOWN_ERROR", "PERMISSION_DENIED",

		// User status messages
		"REGISTERED_EMAIL", "USER_NOT_FOUND", "INCORRECT_PASSWORD", "REGISTERED_EMAIL",

		// API status messages
		"VALIDATION_ERROR",
	}[r]
}
