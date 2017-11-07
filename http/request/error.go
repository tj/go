package request

import "net/http"

// Error is an HTTP error.
type Error int

// Error implementation.
func (e Error) Error() string {
	return http.StatusText(int(e))
}

// IsNotFound returns true if err is a 404.
func IsNotFound(err error) bool {
	e, ok := err.(Error)
	return ok && e == 404
}
