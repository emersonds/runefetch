package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Name    string   `json:"name"`
	Mode    string   `json:"mode"`
	Logo    string   `json:"logo"`
	Modules []string `json:"modules"`
}

// Looks for config file and returns its contents
func GetConfig(confPath string) *Config {
	file, err := os.Open(confPath)
	if err != nil {
		fmt.Printf("Error loading config file: %v\n", err)
		return nil
	}
	defer file.Close() // Close file at end of function

	decoder := json.NewDecoder(file)
	config := &Config{}

	if err := decoder.Decode(config); err != nil {
		fmt.Printf("Error decoding config JSON: %v\n", err)
	}

	return config
}
