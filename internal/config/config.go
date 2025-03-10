package config

import (
	"os"
	"encoding/json"
	"fmt"
)

const configFileName = ".gatorconfig.json"

type Config struct {
    db_url string `json:"db_url"`
	current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	fmt.Println("Entering Read() function")
	config := Config{}

	configPath, err := getConfigFilePath()
	if err != nil {
        return config, err
    }
	fmt.Printf("Config File Path retrieved: %s\n", configPath)

	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	fmt.Println("Config File Read")

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (c Config) SetUser(username string) {
	c.current_user_name = username
	write(c)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	

	return homeDir + configFileName, nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	return err
}