package configs

import (
	"encoding/json"
	"os"
	"log"
)

type Config struct {
	DataBaseDsn string `json:"databaseDSN"`
	ServerPort  string `json:"serverPort"`
}

func LoadConfig() (*Config, error) {
	// Use relative path from current working directory or the path specific to Vercel if needed
	configFile, err := os.Open("config/config.json")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
		return nil, err
	}
	defer configFile.Close()

	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Fatalf("Failed to decode config JSON: %v", err)
		return nil, err
	}
	return &config, nil
}
