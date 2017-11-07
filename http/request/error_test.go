package request

import (
	"testing"

	"github.com/tj/assert"
)

func TestError(t *testing.T) {
	assert.Equal(t, "Not Found", Error(404).Error())
	assert.Equal(t, "Bad Request", Error(400).Error())
	assert.True(t, IsNotFound(Error(404)))
	assert.False(t, IsNotFound(Error(500)))
}
