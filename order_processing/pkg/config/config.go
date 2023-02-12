package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// name of json config file
const filename = "config.json"

// name of the .env for development/test environment configs
const envFile = ".env"

const (
	addressKey   = "DB_ADDR"
	userKey      = "DB_USER"
	passwordKey  = "DB_PASS"
	nameKey      = "DB_NAME"
	inventoryKey = "INV_API"
)

// DatabaseSettings contains the configs of the MySQL database that the server connects to
type DatabaseSettings struct {
	Address  string `json:"address"` // database address in ip:port format
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Config contains all the configs this server requires
type Config struct {
	Port         uint              `json:"port"`
	Database     *DatabaseSettings `json:"database"`
	InventoryApi string            `json:"inventory_api"`
	ShipmentMQ   string            `json:"shipment_queue"`
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

	// find the path to .env file
	envPath := filepath.Join(exPath, envFile)

	// load .env file
	godotenv.Load(envPath)
}

func (config *Config) loadFromEnv() {
	address := os.Getenv(addressKey)
	if len(address) > 0 {
		config.Database.Address = address
	}

	user := os.Getenv(userKey)
	if len(user) > 0 {
		config.Database.User = user
	}

	password := os.Getenv(passwordKey)
	if len(password) > 0 {
		config.Database.Password = password
	}

	name := os.Getenv(nameKey)
	if len(name) > 0 {
		config.Database.Name = name
	}

	inventoryApi := os.Getenv(inventoryKey)
	if len(inventoryApi) > 0 {
		config.InventoryApi = inventoryApi
	}
}

func NewConfig() *Config {
	// create a new siteOptions object
	config := Config{}

	// read config.json first
	config.loadFromFile()

	// override configs with values from environment variables
	config.loadFromEnv()

	return &config
}
