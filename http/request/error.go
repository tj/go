package request

import "net/http"

// Error is an HTTP error.
type Error int

// Error implementation.
func (e Error) Error() string {
	return http.StatusText(int(e))
}

// IsStatus returns true if err is status code.
func IsStatus(err error, code int) bool {
	e, ok := err.(Error)
	return ok && int(e) == code
}

// IsNotFound returns true if err is a 404.
func IsNotFound(err error) bool {
	return IsStatus(err, 404)
}
