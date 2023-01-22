package main

import (
	"encoding/json"
	"log"
	"os"
)

var Config CVEDashConfig

type CVEDashConfig struct {
	NVDAPIKey  string `json:"nvdApiKey"`
	DBFilePath string `json:"dbFilePath"`
}

func LoadConfig(path string) CVEDashConfig {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Could not read config file at "+path+":", err)
	}

	var jsonConfig CVEDashConfig
	json.Unmarshal(contents, &jsonConfig)

	return jsonConfig
}
