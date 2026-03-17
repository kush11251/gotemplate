package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	Server   Server   `json:"server"`
	Database Database `json:"database"`
}

type Server struct {
	Port string `json:"port"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}