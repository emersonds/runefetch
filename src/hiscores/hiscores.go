package hiscores

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gookit/color"
)

type HiscoreEntry struct {
	Name  string `json:"name"`
	Rank  int    `json:"rank"`
	Level int    `json:"level"`
	XP    int    `json:"xp"`    // Exclusive to Skills
	Score int    `json:"score"` // Exclusive to Activities
}

type HiscoreResponse struct {
	Skills     []HiscoreEntry `json:"skills"`
	Activities []HiscoreEntry `json:"activities"`
}

func (hEntry *HiscoreEntry) PrintEntry(isSkill bool, colors [3]color.RGBColor) string {
	// TODO: Switch statement isSkill
	// if skill print XP. if activity print score.
	// copy print from output, add color, and replace in output loop
	switch isSkill {
	case true:
		return color.Sprintf("<fg=colors[0]%s</> <fg=colors[1]Level</> <fg=colors[2]%d</>, <fg=colors[2]%d</> <fg=colors[1]>XP</>, <fg=colors[1]>Rank</> <fg=colors[2]>%d</>",
			hEntry.Name, hEntry.Level, hEntry.XP, hEntry.Rank)
	case false:
		return fmt.Sprintf("%s%s %sScore %s%d, %sRank %s%d",
			colors[0], hEntry.Name, colors[1], colors[2], hEntry.Score, colors[1], colors[2], hEntry.Rank)
	}
	return ""
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

func GetHiscores(requestHTTP string) (hsResponse HiscoreResponse) {
	// Send GET request to OSRS Hiscores API
	response, err := http.Get(requestHTTP)
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

	if err := json.Unmarshal(body, &hsResponse); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	return hsResponse
}
