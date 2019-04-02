package config

import (
	"encoding/json"
	"os"
	"fmt"
)

type Configuration struct {
	Port string `json:"Port"`
	Database struct {
		URL string `json:"URL"`
		Name string `json:"Name"`
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
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)

	Config = &config
    return &config
}
