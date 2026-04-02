package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runefetch/config"
	"runefetch/hiscores"
	"strings"
)

func main() {
	configPath, err := config.ValidateConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}
	playerData := config.GetConfig(configPath)

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
		switch i {
		case 0:
			fmt.Printf("%s\t%s%s%s\n", logoScanner.Text(), config.DefaultAccentColor, playerData.Name, config.ResetColor)
			continue
		case 1:
			fmt.Printf("%s\t%s%s%s\n", logoScanner.Text(), config.DefaultAccentColor, playerData.Mode, config.ResetColor)
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
