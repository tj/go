package array

import (
	"encoding/json"
	"testing"

	"github.com/tj/assert"
)

func TestArray_MarshalJSON(t *testing.T) {
	s := Array{"bar", 123}

	b, err := json.Marshal(s)
	assert.NoError(t, err, "marshal")

	assert.Equal(t, `["bar",123]`, string(b))
}

func TestArray_Empty(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		s := Array{}
		assert.True(t, s.Empty())
	})

	t.Run("when populated", func(t *testing.T) {
		s := Array{"foo"}
		assert.False(t, s.Empty())
	})
}
