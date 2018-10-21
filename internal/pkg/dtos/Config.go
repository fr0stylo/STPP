package dtos

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Server   string
	Database string
	Issuer   string
	Audience string
}


func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
