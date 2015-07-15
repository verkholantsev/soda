package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func GetConfig() Config {
	configPath := os.Getenv("CONFIG")
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
