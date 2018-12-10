package config

import (
	"net/http"
	. "time-logger/internal/pkg/dtos"
	. "time-logger/shared/http-wrappers"
)

type IConfigReader interface {
	GetConfig(w http.ResponseWriter, r *http.Request)
}

type ConfigReader struct {
	Config IConfig
}

func (cr *ConfigReader) GetConfig(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	config := cr.Config.Read()

	RespondWithJson(w, http.StatusOK, config)
}
