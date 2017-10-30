// Package env provides environment variable utilities.
package env

import (
	"fmt"
	"os"
)

// Get panics if the environment variable is missing.
func Get(name string) string {
	if s := os.Getenv(name); s == "" {
		panic(fmt.Errorf("environment variable %q is required", name))
	} else {
		return s
	}
}

// GetDefault returns `value` if environment variable `name` is not present.
func GetDefault(name string, value string) string {
	if s := os.Getenv(name); s == "" {
		return value
	} else {
		return s
	}
}
