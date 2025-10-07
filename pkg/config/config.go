package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Jwt     string `json:"jwt_dew"`
	DBpatch string `json:"db_patch"`
	Port    string `json:"port"`
}

func LoadConfig(patch string) (*Config, error) {
	file, err := os.Open(patch)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var config Config
	decoder := json.NewDecoder(file)
	er := decoder.Decode(&config)
	if er != nil {
		return nil, er
	}
	if envJWT := os.Getenv("JWT_SECRET"); envJWT != "" {
		config.Jwt = envJWT
	}

	if envDBPath := os.Getenv("DB_PATH"); envDBPath != "" {
		config.DBpatch = envDBPath
	}

	if envPort := os.Getenv("PORT"); envPort != "" {
		config.Port = envPort
	}
	return &config, nil
}
