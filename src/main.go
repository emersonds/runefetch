package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runefetch/config"
	"runefetch/hiscores"
	"strings"

	"github.com/gookit/color"
)

func main() {
	configPath, err := config.ValidateConfigDir()
	if err != nil {
		fmt.Printf("%v", err)
	}
	playerData := config.GetConfig(configPath)
	var playerColors [3]color.RGBColor = config.GetColors(*playerData)

	var playerHiscores string
	if playerData != nil {
		playerHiscores = hiscores.HiscoresBuilder(playerData.Name, playerData.Mode)
	} else {
		fmt.Printf("Invalid player config\n")
		return
	}

	// Unmarshal json into HiscoreResponse struct.
	// Makes output easier to read and use.
	hiscoreData := hiscores.GetHiscores(playerHiscores)

	// Only display hiscores from the config
	var displaySkills []hiscores.HiscoreEntry
	var displayActivities []hiscores.HiscoreEntry
	for _, module := range playerData.Modules {
		for _, skill := range hiscoreData.Skills {
			if strings.EqualFold(skill.Name, module) {
				displaySkills = append(displaySkills, skill)
			}
		}
		for _, activity := range hiscoreData.Activities {
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
		fmt.Printf("%s\t", logoScanner.Text())
		switch i {
		case 0:
			c := playerColors[0].Sprint(playerData.Name)
			color.Print(c + "\n")
			continue
		case 1:
			c := playerColors[0].Sprint(playerData.Mode)
			color.Print(c + "\n")
			continue
		}
		if skillsCount < len(displaySkills) {
			fmt.Printf("%s\n", displaySkills[skillsCount].PrintEntry(true, playerColors))
			skillsCount++
		} else if activitiesCount < len(displayActivities) {
			fmt.Printf("%s\n", displayActivities[activitiesCount].PrintEntry(false, playerColors))
			activitiesCount++
		} else {
			// Needed so each line of logo goes on a new line when no other text is printed
			fmt.Print("\n")
		}
	}
}
