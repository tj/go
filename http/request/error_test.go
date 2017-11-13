package request

import (
	"testing"

	"github.com/tj/assert"
)

func TestError(t *testing.T) {
	assert.Equal(t, "Not Found", Error(404).Error())
	assert.Equal(t, "Bad Request", Error(400).Error())
}

func TestIsNotFound(t *testing.T) {
	assert.True(t, IsNotFound(Error(404)))
	assert.False(t, IsNotFound(Error(500)))
}

func TestIsStatus(t *testing.T) {
	assert.True(t, IsStatus(Error(404), 404))
	assert.True(t, IsStatus(Error(500), 500))
	assert.False(t, IsStatus(Error(500), 400))
}

func TestIsClient(t *testing.T) {
	assert.True(t, IsClient(Error(404)))
	assert.False(t, IsClient(Error(500)))
}

func TestIsServer(t *testing.T) {
	assert.False(t, IsServer(Error(404)))
	assert.True(t, IsServer(Error(500)))
}
