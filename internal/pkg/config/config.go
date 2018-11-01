package config

import (
	"encoding/json"
	"log"
	"net/http"

	. "time-logger/internal/pkg/dtos"
)

type IConfigReader interface {
	GetConfig()
}

type Reader struct {
}


func (r *Reader) GetConfig() (Config) {
	var config Config
	response, err := http.Get("http://config-service:3000/")

	defer response.Body.Close()
	if json.NewDecoder(response.Body).Decode(&config); err != nil {
		log.Fatal(err.Error())
	}

	return config
}
