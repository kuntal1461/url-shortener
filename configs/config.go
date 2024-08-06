package configs

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DataBaseDsn string `json:"databaseDSN"`
	ServerPort  string `json:"serverPort"`
}

func LoadConfig() (*Config, error) {
	// Open the config file for default configurations
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

	// Use environment variables if set
	if dsn := os.Getenv("POSTGRES_URL"); dsn != "" {
		config.DataBaseDsn = dsn
	} else {
		log.Println("Using default database DSN from config file")
	}

	return &config, nil
}
