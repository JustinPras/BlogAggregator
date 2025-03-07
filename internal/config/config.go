package config

import (
	"os"
	"fmt"
	"json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
    db_url string `json:"db_url"`
	current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	config := Config{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}

	filePath := homeDir + configFileName

	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(filePath, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (c Config) SetUser(username string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}

	config := Config{}
	filePath := homeDir + configFileName


	
}

// func getConfigFilePath() (string, error) {

// }

func write(cfg Config) error {
	
}