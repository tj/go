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

// IsClient returns true if err represents a 4xx error.
func IsClient(err error) bool {
	e, ok := err.(Error)
	return ok && e >= 400 && e < 500
}

// IsServer returns true if err represents a 5xx error.
func IsServer(err error) bool {
	e, ok := err.(Error)
	return ok && e >= 500
}

// IsNotFound returns true if err is a 404.
func IsNotFound(err error) bool {
	return IsStatus(err, 404)
}

// Param returns the parameter by name.
func Param(r *http.Request, name string) string {
	return r.URL.Query().Get(name)
}

// ParamDefault returns the parameter by name or default value.
func ParamDefault(r *http.Request, name, value string) string {
	if s := Param(r, name); s != "" {
		return s
	}
	return value
}
