package config

import (
	"encoding/json"
	"log"
	"os"
)

var Config CVEDashConfig

type CVEDashConfig struct {
	NVDAPIKey                string `json:"nvdApiKey"`
	DBFilePath               string `json:"dbFilePath"`
	MinutesBetweenNVDUpdates uint64 `json:"minutesBetweenNVDUpdates"`
	APIRateLimitPerMinute    uint64 `json:"apiRateLimitPerMinute"`
}

func LoadConfig(path string) CVEDashConfig {
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Could not read config file at "+path+":", err)
	}

	jsonConfig := CVEDashConfig{ // default values
		NVDAPIKey:                "",
		DBFilePath:               "cvedash-data.db",
		MinutesBetweenNVDUpdates: 120,
		APIRateLimitPerMinute:    250,
	}
	json.Unmarshal(contents, &jsonConfig)

	return jsonConfig
}
