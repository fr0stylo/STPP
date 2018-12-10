package http_wrappers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	response := httptest.NewRecorder()

	RespondWithError(response, http.StatusConflict, "Error Occurred")

	assert.Equal(t, "{\"error\":\"Error Occurred\"}", response.Body.String())
	assert.Equal(t, http.StatusConflict, response.Code)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, response.Header())
}

func TestRespondWithJson(t *testing.T) {
	response := httptest.NewRecorder()

	payload := map[string]string{"foo": "bar"}

	RespondWithJson(response, http.StatusOK, payload)

	assert.Equal(t, "{\"foo\":\"bar\"}", response.Body.String())
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, http.Header{"Content-Type": []string{"application/json"}}, response.Header())

}
