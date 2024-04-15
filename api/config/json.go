package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func GetURL() (string, error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Error reading config file: ", err)
		return "", err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
		return "", err
	}

	return config.Ip + ":" + config.Port, nil
}
