package main

import (
	"fmt"
	//"os"
	//"net/http"
	//"encoding/json"
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

type RuneScapePlayer struct {
	name   string
	mode   string   // Normal, Ironman, etc.
	skills []string // Attack, Defence, etc.
}

// Builds the full http URL for Old School Hiscores API
func HiscoresBuilder(name string, mode string) string {
	return hiscoresHttpPrefix + mode + hiscoresHttpSuffix + name
}

func GetPlayerStats(playerName string) {

}

func main() {
	fmt.Println(HiscoresBuilder(testName, normalMode))
}
