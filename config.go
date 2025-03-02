package main

import (
	Logger "api-spammer/logger"
	"encoding/json"
	"os"
)

type Endpoint struct {
	Url    string                 `json:"url"`
	Method string                 `json:"method"`
	Data   map[string]interface{} `json:"data"`
}
type Config struct {
	Endpoints     []Endpoint `json:"endpoints"`
	Clients       int        `json:"clients"`
	TotalRequests int        `json:"total_requests"`
	WithProxy     bool       `json:"with_proxy"`
}

func LoadConfig() Config {
	file, err := os.Open("./config.json")
	if err != nil {
		Logger.Log(Logger.ColorError, "Failed to open config file")
		os.Exit(1)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		Logger.Log(Logger.ColorError, "Failed to parse config file", err)
		os.Exit(1)
	}

	return config
}
