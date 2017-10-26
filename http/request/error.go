package request

import "net/http"

// Error is an HTTP error.
type Error struct {
	Status int
}

// Error implementation.
func (e Error) Error() string {
	return http.StatusText(e.Status)
}
