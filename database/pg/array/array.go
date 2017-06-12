// Package array implements a JSONB array.
package array

import (
	"database/sql/driver"
	"encoding/json"
)

// Array type.
type Array []interface{}

// New returns an empty array.
func New() Array {
	return make(Array, 0)
}

// Scan implementation.
func (v *Array) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		if err := json.Unmarshal(src.([]byte), &v); err != nil {
			return err
		}
		return nil
	default:
		return nil
	}
}

// Value implementation.
func (v Array) Value() (driver.Value, error) {
	if v.Empty() {
		return "[]", nil
	}

	b, err := json.Marshal(v)
	return string(b), err
}

// Empty checks if the set is empty.
func (v Array) Empty() bool {
	return len(v) == 0
}

// Interface assertion.
var _ driver.Value = (Array)(nil)
