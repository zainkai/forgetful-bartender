package config

import (
	"encoding/json"
	"os"

	"github.com/zainkai/forgetful-bartender/pkg/logger"
)

type Configuration struct {
	Port     string `json:"Port"`
	Database struct {
		URL        string `json:"URL"`
		Name       string `json:"Name"`
		Collection string `json:"Collection"`
	} `json:"Database"`
}

var Config *Configuration

func LoadConfig() *Configuration {
	if Config != nil { // dont load config multiple times
		return Config
	}
	var config Configuration
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		logger.Err("configs/config", "couldn't open config.json", err)
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	Config = &config
	return &config
}
