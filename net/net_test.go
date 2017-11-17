package net

import (
	"testing"

	"github.com/tj/assert"
)

func TestGetCert(t *testing.T) {
	t.Run("valid https", func(t *testing.T) {
		c, err := GetCert("https://apex.sh")
		assert.NoError(t, err, "cert")
		assert.NotEmpty(t, c, "empty cert")
	})

	t.Run("explicit port", func(t *testing.T) {
		c, err := GetCert("https://apex.sh:443")
		assert.NoError(t, err, "cert")
		assert.NotEmpty(t, c, "empty cert")
	})

	t.Run("http", func(t *testing.T) {
		_, err := GetCert("http://apex.sh")
		assert.EqualError(t, err, `https only`)
	})
}

func TestGetCertSummary(t *testing.T) {
	t.Run("valid https", func(t *testing.T) {
		c, err := GetCertSummary("https://apex.sh")
		assert.NoError(t, err, "cert")
		assert.NotEmpty(t, c.IssuedAt)
		assert.NotEmpty(t, c.ExpiresAt)
		assert.Equal(t, "Amazon", c.Issuer.Name)
		assert.Equal(t, "Amazon", c.Issuer.Organization)
		assert.Equal(t, "US", c.Issuer.Country)
		assert.True(t, c.IssuedAt.Before(c.ExpiresAt))
	})
}
