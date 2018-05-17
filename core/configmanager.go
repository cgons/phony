package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// ParseConfig loads and parses a config file specified (absolute path)
func ParseConfig(configPath string) []Route {
	var config []Route

	configPath = GetConfigPath(configPath)
	configFile := LoadConfig(configPath)
	log.Printf("Loaded config file: %s", configPath)

	json.Unmarshal(configFile, &config) // parse + dump JSON into config map
	return config
}

// LoadConfig loads a phonyconfig.json file from disk.
func LoadConfig(configPath string) []byte {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// GetConfigPath returns the absolute path to the config file.
// ...simply returning the absolute path provided or, if path is an empty string
// the absolute path to phonyconfig.json - default location (cur. working dir.)
func GetConfigPath(path string) string {
	if path == "" {
		path = GetWorkingDir() + "/phonyconfig.json"
	}
	return path
}

// GetWorkingDir returns the absolute path to the current working dir.
// (ie. the dir from which this executable was run).
func GetWorkingDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
