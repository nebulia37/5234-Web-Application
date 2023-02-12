package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

// name of json config file
const filename = "config.json"

// Config contains all the configs this server requires
type Config struct {
	Port       uint   `json:"port"`
	ShipmentMQ string `json:"shipment_queue"`
}

func (config *Config) loadFromFile() {
	// find the path to current executable
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	// open config.json file
	file, err := os.Open(filepath.Join(exPath, filename))
	if err != nil {
		log.Printf("Failed to open config file [%s]. Error[%v]\n", filename, err)
		panic(err)
	}

	defer file.Close()

	bytes, _ := io.ReadAll(file)

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Printf("Failed to parse config file [%s]. Error[%v]\n", filename, err)
		panic(err)
	}
}

func NewConfig() *Config {
	// create a new siteOptions object
	config := Config{}

	// read config.json first
	config.loadFromFile()

	return &config
}
