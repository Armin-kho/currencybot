package db

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token       string `json:"token"`
	DBPath      string `json:"db_path"`
	Language    string `json:"language"`
	NavasanKey  string `json:"navasan_key"`
}

func LoadConfig() (*Config, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
