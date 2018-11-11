package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time-logger/internal/pkg/dtos"
	"time-logger/internal/pkg/http-wrappers"
)

var httpClientMock *http_wrappers.HttpClientMock

func setUp(item interface{}, errr error, status int) {
	httpClientMock = &http_wrappers.HttpClientMock{
		GetFunc: func(url string) (*http.Response, error) {
			res, _ := json.Marshal(item)
			return &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader(res)),
				StatusCode: status,
			}, errr
		},
	}
}

func setUpString(json string) {
	httpClientMock = &http_wrappers.HttpClientMock{
		GetFunc: func(url string) (*http.Response, error) {
			return &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(json)),
				StatusCode: http.StatusOK,
			}, nil
		},
	}
}

func TestNewReader(t *testing.T) {
	reader := NewReader(httpClientMock)

	assert.IsType(t, &Reader{}, reader)

}

func TestReader_GetConfig(t *testing.T) {
	config := dtos.Config{
		Server:   "server",
		Audience: "aud",
		Database: "db",
		Issuer:   "iss",
	}

	setUp(config, nil, http.StatusOK)

	reader := NewReader(httpClientMock)

	res, err := reader.GetConfig()

	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, config, res)
}

func TestReader_GetConfig2(t *testing.T) {
	config := dtos.TimeEntryDTO{Description: "as"}

	setUp(config, nil, http.StatusOK)

	reader := NewReader(httpClientMock)

	res, _ := reader.GetConfig()

	assert.IsType(t, dtos.Config{}, res)
	assert.Equal(t, dtos.Config{}, res)
}

func TestReader_GetConfig_ShouldReturnError(t *testing.T) {
	config := dtos.Config{
		Server:   "server",
		Audience: "aud",
		Database: "db",
		Issuer:   "iss",
	}

	setUp(config, fmt.Errorf("%s", "Error"), http.StatusBadRequest)

	reader := NewReader(httpClientMock)

	_, err := reader.GetConfig()

	assert.Equal(t, "Error", err.Error())
}
