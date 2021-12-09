package http

// Error is used to pass an error during the request
type Error struct {
	Err    error
	Status int
}

// ErrorResponse is the form used for API responses in the API.
type ErrorResponse struct {
	Error string
}

// NewResponseError wraps a provided error with an HTTP status code.
func NewResponseError(err error, status int) error {
	return &Error{
		Err:    err,
		Status: status,
	}
}

// Error implements the error interface
func (err *Error) Error() string {
	return err.Err.Error()
}
