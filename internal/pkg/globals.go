package pkg

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

var globalsMutex sync.Mutex

type Globals struct {
	Total int `json:"total"`
	Today int `json:"today"`
}

func GetGlobals() (*Globals, error) {
	globalsMutex.Lock()
	defer globalsMutex.Unlock()

	path := os.Getenv("GLOBALS_JSON_FILEPATH")
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var globals Globals
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteValue, &globals)
	if err != nil {
		return nil, err
	}

	return &globals, nil
}

func WriteGlobals(globals *Globals) error {
	globalsMutex.Lock()
	defer globalsMutex.Unlock()

	jsonData, err := json.Marshal(globals)
	if err != nil {
		return err
	}

	path := os.Getenv("GLOBALS_JSON_FILEPATH")
	file, err := os.OpenFile(path, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func ResetToday() error {
	globals, err := GetGlobals()
	if err != nil {
		return err
	}

	globals.Today = 0
	err = WriteGlobals(globals)
	if err != nil {
		return err
	}

	return nil
}
