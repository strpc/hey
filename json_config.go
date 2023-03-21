package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type config struct {
	Count       int       `json:"count,omitempty"`
	Concurrency int       `json:"concurrency,omitempty"`
	Timeout     int       `json:"timeout,omitempty"`
	Requests    []request `json:"requests"`
}

type request struct {
	Method   string                 `json:"method"`
	Url      string                 `json:"url"`
	Headers  map[string]string      `json:"headers,omitempty"`
	Body     map[string]interface{} `json:"body,omitempty"`
	Cookies  map[string]string      `json:"cookies,omitempty"`
	Username string                 `json:"username,omitempty"`
	Password string                 `json:"password,omitempty"`
	Host     string                 `json:"host,omitempty"`
}

func parseJsonConfig(configPath string) config {
	file, err := os.Open(configPath)
	if err != nil {
		usageAndExit(err.Error())
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var config config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		usageAndExit("invalid json config")
	}
	return config
}
