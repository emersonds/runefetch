package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	// These match the colors set in your terminal with your color scheme.
	defaultAccentColor = "\033[0;35m" // Magenta
	//defaultSecondaryColor = "\033[0;32m" // Green
	resetColor = "\033[0m"
)

type Config struct {
	Name    string   `json:"name"`
	Mode    string   `json:"mode"`
	Logo    string   `json:"logo"`
	Modules []string `json:"modules"`
}

type HiscoreEntry struct {
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Level int    `json:"level"`
	XP    int    `json:"xp"`    // Exclusive to Skills
	Score int    `json:"score"` // Exclusive to Activities
}

func (hEntry *HiscoreEntry) PrintEntry(isSkill bool) string {
	// TODO: Switch statement isSkill
	// if skill print XP. if activity print score.
	// copy print from output, add color, and replace in output loop
	switch isSkill {
	case true:
		return fmt.Sprintf("%s%s %sLevel %d, %d XP, Rank %d",
			defaultAccentColor, hEntry.Name, resetColor, hEntry.Level, hEntry.XP, hEntry.Rank)
	case false:
		return fmt.Sprintf("%s%s %sScore %d, Rank %d",
			defaultAccentColor, hEntry.Name, resetColor, hEntry.Score, hEntry.Rank)
	}
	return ""
}

type HiscoreResponse struct {
	Skills     []HiscoreEntry `json:"skills"`
	Activities []HiscoreEntry `json:"activities"`
}

// Builds the full http URL for Old School Hiscores API
func HiscoresBuilder(name string, mode string) string {
	mode = strings.ToLower(mode)
	var hiscoreHTTP string

	switch mode {
	case "normal", "main":
		hiscoreHTTP = "https://secure.runescape.com/m=hiscore_oldschool/index_lite.json?player="
	case "iron", "ironman":
		hiscoreHTTP = "https://secure.runescape.com/m=hiscore_oldschool_ironman/index_lite.json?player="
	case "hc", "hardcore", "hardcore iron", "hardcore ironman":
		hiscoreHTTP = "https://secure.runescape.com/m=hiscore_oldschool_hardcore_ironman/index_lite.json?player="
	case "ultimate", "ultimate iron", "ultimate ironman":
		hiscoreHTTP = "https://secure.runescape.com/m=hiscore_oldschool_ultimate/index_lite.json?player="
	}

	return hiscoreHTTP + name
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

func main() {
	confDir, dirErr := os.UserConfigDir()

	var configPath string

	if dirErr == nil {
		configPath = filepath.Join(confDir, "runefetch", "config.json")
	} else {
		fmt.Printf("Unable to locate config directory: %v", dirErr)
		return
	}

	playerData := GetConfig(configPath)

	var playerHiscores string
	if playerData != nil {
		playerHiscores = HiscoresBuilder(playerData.Name, playerData.Mode)
	} else {
		fmt.Printf("Invalid player config\n")
		return
	}

	// Send GET request to OSRS Hiscores API
	response, err := http.Get(playerHiscores)
	if err != nil {
		fmt.Printf("GET Request Error: %v\n", err)
		return
	}
	defer response.Body.Close() // Close body at end of function

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\nStatus: %v\n", response.StatusCode, response.Status)
		return
	}

	// Converts response body to []byte for json unmarshaling
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Unmarshal json into HiscoreResponse struct.
	// Makes output easier to read and use.
	var hiscore HiscoreResponse
	if err := json.Unmarshal(body, &hiscore); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Only display hiscores from the config
	var displaySkills []HiscoreEntry
	var displayActivities []HiscoreEntry
	for _, module := range playerData.Modules {
		for _, skill := range hiscore.Skills {
			if strings.EqualFold(skill.Name, module) {
				displaySkills = append(displaySkills, skill)
			}
		}
		for _, activity := range hiscore.Activities {
			if strings.EqualFold(activity.Name, module) {
				displayActivities = append(displayActivities, activity)
			}
		}
	}
	//fmt.Println("Display Modules:",displayActivities,displaySkills)

	// Get logo to display from config
	logoPath := filepath.Join("logos", string(playerData.Logo+".txt"))
	logoData, err := os.Open(logoPath)
	if err != nil {
		fmt.Printf("Error loading logo: %v\n", err)
	}
	defer logoData.Close() // Close file at end of function
	logoScanner := bufio.NewScanner(logoData)

	// Output display. Each loop prints a line from the logo and a skill/activity
	skillsCount := 0
	activitiesCount := 0
	for i := 0; logoScanner.Scan(); i++ {
		switch i {
		case 0:
			fmt.Printf("%s\t%s%s%s\n", logoScanner.Text(), defaultAccentColor, playerData.Name, resetColor)
			continue
		case 1:
			fmt.Printf("%s\t%s%s%s\n", logoScanner.Text(), defaultAccentColor, playerData.Mode, resetColor)
			continue
		}
		if skillsCount < len(displaySkills) {
			fmt.Printf("%s\t%s\n", logoScanner.Text(), displaySkills[skillsCount].PrintEntry(true))
			skillsCount++
		} else if activitiesCount < len(displayActivities) {
			fmt.Printf("%s\t%s\n", logoScanner.Text(), displayActivities[activitiesCount].PrintEntry(false))
			activitiesCount++
		} else {
			fmt.Printf("%s\n", logoScanner.Text())
		}
	}
}
