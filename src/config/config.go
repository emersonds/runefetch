package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gookit/color"
)

var defaultColors = [3]color.Basic{color.Magenta, color.Green, color.Blue}

type Config struct {
	Name    string   `json:"name"`
	Mode    string   `json:"mode"`
	Logo    string   `json:"logo"`
	Colors  []string `json:"colors"`
	Modules []string `json:"modules"`
}

// Verifies the config file exists within the correct config directory
func ValidateConfigDir() (string, error) {
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

// Verifies config colors, returning default values if invalid
func GetColors(conf Config) (colors [3]color.RGBColor) {
	for i, confColor := range conf.Colors {
		switch len(confColor) {
		case 6, 7:
			colors[i] = color.HEX(confColor)
		case 1, 2, 3:
			s, err := strconv.Atoi(confColor)
			// confColor contains letters, fallback to HEX
			if err != nil {
				fmt.Printf("Error parsing color: %v\n", err)
				colors[i] = color.HEX(confColor)
			}
			colors[i] = color.RGBColor(color.C256ToRgb(uint8(s)))
		default:
			colors[i] = defaultColors[i].RGB()
		}
	}

	//fmt.Printf("Colors: %v\n", colors)
	return colors
}
