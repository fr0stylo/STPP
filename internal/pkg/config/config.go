package config

import (
	"encoding/json"
	. "time-logger/internal/pkg/dtos"
	"time-logger/shared/http-wrappers"
)

type IConfigReader interface {
	GetConfig()
}

type Reader struct {
	http http_wrappers.HttpClient
}

func NewReader(client http_wrappers.HttpClient) *Reader {
	return &Reader{http: client}
}

func (r *Reader) GetConfig() (Config, error) {
	var config Config
	response, err := r.http.Get("http://config-service:3000/")
	defer func() {
		if response.Body != nil {
			response.Body.Close()
		}
	}()

	if err != nil {
		return config, err
	}

	json.NewDecoder(response.Body).Decode(&config)

	return config, nil
}
