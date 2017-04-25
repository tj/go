package set

import (
	"encoding/json"
	"testing"

	"github.com/tj/assert"
)

func TestSet_MarshalJSON(t *testing.T) {
	s := Set{"bar"}

	b, err := json.Marshal(s)
	assert.NoError(t, err, "marshal")

	assert.Equal(t, `["bar"]`, string(b))
}

func TestSet_Empty(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		s := Set{}
		assert.True(t, s.Empty())
	})

	t.Run("when populated", func(t *testing.T) {
		s := Set{"foo"}
		assert.False(t, s.Empty())
	})
}

func TestSet_Has(t *testing.T) {
	s := Set{"foo", "bar"}
	assert.True(t, s.Has("foo"))
	assert.True(t, s.Has("bar"))
	assert.False(t, s.Has("baz"))
}

func TestSet_Remove(t *testing.T) {
	s := Set{"foo", "bar", "baz"}

	s.Remove("bar")
	assert.Equal(t, Set{"foo", "baz"}, s)

	s.Remove("bar")
	assert.Equal(t, Set{"foo", "baz"}, s)

	s.Remove("foo")
	assert.Equal(t, Set{"baz"}, s)

	s.Remove("baz")
	s.Remove("something")
	s.Remove("")
	assert.Equal(t, Set{}, s)
}

func TestSet_Value(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		s := Set{}
		v, err := s.Value()
		assert.NoError(t, err, "value")
		assert.Equal(t, `{}`, v)
	})

	t.Run("when populated", func(t *testing.T) {
		s := Set{"foo"}
		v, err := s.Value()
		assert.NoError(t, err, "value")
		assert.Equal(t, `{"foo":true}`, v)
	})
}
