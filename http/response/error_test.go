package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tj/assert"
)

func TestError(t *testing.T) {
	res := httptest.NewRecorder()
	Error(res, http.StatusBadRequest)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "Bad Request\n", string(res.Body.Bytes()))
	assert.Equal(t, "text/plain; charset=utf-8", res.HeaderMap["Content-Type"][0])
}
