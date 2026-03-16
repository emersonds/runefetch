package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Used for http builder
const (
	testName           = "B0aty"
	hiscoresHttpPrefix = "https://secure.runescape.com/m=hiscore_oldschool"
	hiscoresHttpSuffix = "/index_lite.json?player="
	normalMode         = ""
	ironmanMode        = "_ironman"
	hardcoreMode       = "_hardcore_ironman"
	ultimateMode       = "_ultimate"
)

type RuneScapePlayer struct {
	name string
	mode string // Normal, Ironman, etc.
}

type HiscoreEntry struct {
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Level int    `json:"level"`
	XP    int    `json:"xp"`    // Exclusive to Skills
	Score int    `json:"score"` // Exslusive to Activities
}

type HiscoreResponse struct {
	Skills     []HiscoreEntry `json:"skills"`
	Activities []HiscoreEntry `json:"activities"`
}

// Builds the full http URL for Old School Hiscores API
func HiscoresBuilder(name string, mode string) string {
	return hiscoresHttpPrefix + mode + hiscoresHttpSuffix + name
}

func main() {
	// testJson := []byte(``)

	player := RuneScapePlayer{name: testName, mode: normalMode}

	// Builds test http request using B0aty
	// https://secure.runescape.com/m=hiscore_oldschool/index_lite.json?player=B0aty
	playerHiscores := HiscoresBuilder(player.name, player.mode)

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

	// Output, print results of get request (skills/activities)
	fmt.Println("===== B0aty Scores =====")
	for _, skills := range hiscore.Skills {
		fmt.Printf("Skill: %s | Rank: %d | Level: %d | XP: %d\n",
			skills.Name, skills.Rank, skills.Level, skills.XP)
	}
	for _, activities := range hiscore.Activities {
		fmt.Printf("Activity: %s | Rank: %d | Score: %d\n",
			activities.Name, activities.Rank, activities.Score)
	}
}
