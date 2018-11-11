package dtos

import (
	"github.com/BurntSushi/toml"
	"log"
)

//go:generate moq -out IConfig_mock.go . IConfig
type IConfig interface {
	Read() Config
}

type Config struct {
	Server   string
	Database string
	Issuer   string
	Audience string
}

func (Config) Read() Config {
	c := Config{}
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}

	return c
}
