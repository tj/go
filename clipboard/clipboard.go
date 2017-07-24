package clipboard

import "github.com/atotto/clipboard"

// Write to clipboard.
func Write(s string) error {
	return clipboard.WriteAll(s)
}

// Read from clipboard.
func Read() (string, error) {
	return clipboard.ReadAll()
}
