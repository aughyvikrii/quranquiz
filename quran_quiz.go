package quranquiz

import (
	"embed"
	"encoding/json"
)

//go:embed quran.json
var fileJSON embed.FS

var QURAN Quran

func init() {
	file, _ := fileJSON.ReadFile("quran.json")
	_ = json.Unmarshal([]byte(file), &QURAN)
}
