package config

import (
	"encoding/json"
	"os"
)

type LanguageConfig struct {
	Command string
}

// Config map for language
var jsonMap = make(map[string]LanguageConfig)

func LoadLanguageCommand() (err error) {
	var directory string
	// JSON file byte slice
	var jsonByte []byte
	// JSON golang map

	directory, err = os.Getwd()
	if err != nil {
		return
	}

	jsonByte, err = os.ReadFile(directory + "/internal/config/config.json")
	if err != nil {
		return
	}

	// Unmarshal byte data to golang struct
	err = json.Unmarshal(jsonByte, &jsonMap)
	return
}
