package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	WeaviateHost string `json:"weaviateHost"`
	WeaviatePort int    `json:"weaviatePort"`
}

func ReadConfig(path string) (*Config, error) {
	var config Config

	bytes, err := os.ReadFile(path)
	if err != nil {
		return &config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return &config, err
	}

	return &config, nil
}
