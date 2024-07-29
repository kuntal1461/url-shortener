package configs

import (
	"encoding/json"
	"os"
)



type Config struct{

	DataBaseDsn string  `json:"databaseDSN"`
	ServerPort string 	`json:"serverPort"`
}

func LoadConfig()(*Config, error){
	configFile,err :=os.Open("../config/config.json")
	if err !=nil {
		return nil, err
	}
	defer configFile.Close() 

	var config Config
	jsonParser := json.NewDecoder(configFile)
	if err=jsonParser.Decode(&config);err !=nil {
		return nil ,err
	}
	return &config ,nil
}