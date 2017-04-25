// Package set implements a JSONB set. The Go-land type is backed by a string slice,
// while the Postgres JSONB value is backed by an object.
package set

import (
	"database/sql/driver"
	"encoding/json"
)

// Set type.
type Set []string

// New returns an empty set.
func New() Set {
	return make(Set, 0)
}

// Scan implementation.
func (v *Set) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		var m map[string]bool

		if err := json.Unmarshal(src.([]byte), &m); err != nil {
			return err
		}

		for k := range m {
			*v = append(*v, k)
		}

		return nil
	default:
		return nil
	}
}

// Value implementation.
func (v Set) Value() (driver.Value, error) {
	if v.Empty() {
		return "{}", nil
	}

	m := make(map[string]bool)

	for _, s := range v {
		m[s] = true
	}

	b, err := json.Marshal(m)
	return string(b), err
}

// Add value to the set.
func (v *Set) Add(value string) {
	if !v.Has(value) {
		*v = append(*v, value)
	}
}

// Remove value from the set.
func (v *Set) Remove(value string) {
	for i, s := range *v {
		if s == value {
			*v = append((*v)[:i], (*v)[i+1:]...)
		}
	}
}

// Has returns true if the value is present.
func (v Set) Has(value string) bool {
	for _, s := range v {
		if s == value {
			return true
		}
	}
	return false
}

// Values returns the set values as a slice.
func (v Set) Values() []string {
	return v
}

// Empty checks if the set is empty.
func (v Set) Empty() bool {
	return len(v) == 0
}

// Interface assertion.
var _ driver.Value = (Set)(nil)
