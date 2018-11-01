package config

import (
	"net/http"
	. "time-logger/internal/pkg/dtos"
	. "time-logger/internal/pkg/http-wrappers"
)

var config = Config{}

type IConfigReader interface {
	GetConfig(w http.ResponseWriter, r *http.Request)
} 

type ConfigReader struct {
	
}

func (cr *ConfigReader) GetConfig(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	config.Read()

	RespondWithJson(w, http.StatusOK, config)
}
