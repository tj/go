// Package object implements a JSONB object.
package object

import (
	"database/sql/driver"
	"encoding/json"
)

// Object type.
type Object map[string]interface{}

// Scan implements the Scanner interface.
func (v *Object) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		return json.Unmarshal(src.([]byte), v)
	default:
		return nil
	}
}

// Value implements the Valuer interface.
func (v Object) Value() (driver.Value, error) {
	if v.Empty() {
		return "{}", nil
	}

	b, err := json.Marshal(v)
	return string(b), err
}

// Keys of the object.
func (v Object) Keys() (keys []string) {
	for k := range v {
		keys = append(keys, k)
	}
	return keys
}

// Size returns the number of values.
func (v *Object) Size() int {
	return len(*v)
}

// Empty checks if the json has values.
func (v *Object) Empty() bool {
	return len(*v) == 0
}

var _ driver.Value = (Object)(nil)
