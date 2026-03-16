package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	testName           = "B0aty"
	hiscoresHttpPrefix = "https://secure.runescape.com/m=hiscore_oldschool"
	hiscoresHttpSuffix = "/index_lite.json?player="
	normalMode         = ""
	ironmanMode        = "_ironman"
	hardcoreMode       = "_hardcore_ironman"
	ultimateMode       = "_ultimate"
)

var test_byte = []byte(`{
	"name": "B0aty",
	"skills": [{
		"id": 0,
		"name": "Overall",
		"rank": 2214,
		"level": 2376,
		"xp": 1203699302
		},
	{
		"id": 1,
		"name": "Attack",
		"rank": 29650,
		"level": 99,
		"xp": 33389881
		}]
	}`)

type RuneScapePlayer struct {
	name string
	mode string // Normal, Ironman, etc.
}

type HiscoreEntry struct {
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Level int    `json:"level"`
	XP    int    `json:"xp"`
}

type HiscoreResponse struct {
	Skills []HiscoreEntry `json:"skills"`
	// Activities map[string][]HiscoreEntry `json:"activities"`
}

// Builds the full http URL for Old School Hiscores API
func HiscoresBuilder(name string, mode string) string {
	return hiscoresHttpPrefix + mode + hiscoresHttpSuffix + name
}

func GetPlayerStats(playerName string) {
}

func main() {
	// testJson := []byte(``)

	player := RuneScapePlayer{name: testName, mode: normalMode}
	playerHiscores := HiscoresBuilder(player.name, player.mode)

	// Send GET request to OSRS Hiscors API
	response, err := http.Get(playerHiscores)
	if err != nil {
		fmt.Printf("GET Request Error: %v\n", err)
	}
	defer response.Body.Close() // Close body at end of function

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Unexpected status code: %d\nStatus: %v\n", response.StatusCode, response.Status)
	}

	// body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
	}

	var hiscore HiscoreResponse
	if err := json.Unmarshal(test_byte, &hiscore); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
	}

	fmt.Println("Test Hiscores")
	for _, skills := range hiscore.Skills {
		fmt.Printf("Skill: %s | Rank: %d | Level: %d | XP: %d\n",
			skills.Name, skills.Rank, skills.Level, skills.XP)
	}
}
