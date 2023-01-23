package main

import (
	"encoding/json"
	"log"
	"os"
)

var Config CVEDashConfig

type CVEDashConfig struct {
	NVDAPIKey             string `json:"nvdApiKey"`
	DBFilePath            string `json:"dbFilePath"`
	MinutesBetweenUpdates uint64 `json:"minutesBetweenUpdates"`
}

func LoadConfig(path string) CVEDashConfig {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Could not read config file at "+path+":", err)
	}

	jsonConfig := CVEDashConfig{ // default values
		NVDAPIKey:             "",
		DBFilePath:            "cvedash-data.db",
		MinutesBetweenUpdates: 10,
	}
	json.Unmarshal(contents, &jsonConfig)

	return jsonConfig
}
