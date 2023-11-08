package quranquiz

import (
	"encoding/json"
	"os"
)

var QURAN Quran

func init() {
	file, _ := os.ReadFile("quran.json")

	_ = json.Unmarshal([]byte(file), &QURAN)
}
