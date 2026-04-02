package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	// These match the colors set in your terminal with your color scheme.
	DefaultAccentColor = "\033[0;35m" // Magenta
	//defaultSecondaryColor = "\033[0;32m" // Green
	ResetColor = "\033[0m"
)

type Config struct {
	Name    string   `json:"name"`
	Mode    string   `json:"mode"`
	Logo    string   `json:"logo"`
	Modules []string `json:"modules"`
}

// Verifies the config file exists within the correct config directory
func ValidateConfig() (string, error) {
	confDir, dirErr := os.UserConfigDir()

	if dirErr == nil {
		return filepath.Join(confDir, "runefetch", "config.json"), nil
	} else {
		return "", fmt.Errorf("Unable to locate config directory: %v", dirErr)
	}
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
