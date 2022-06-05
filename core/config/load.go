package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"deck/core/constant"
)

const defaultLogLevel = "info"

type Config struct {
	Port     string `json:"port"`
	LogLevel string `json:"log_level"`
	PGConfig
}

type PGConfig struct {
	PostgresUser     string `json:"postgres_user"`
	PostgresPassword string `json:"postgres_password"`
	PostgresServer   string `json:"postgres_server"`
	PostgresPort     string `json:"postgres_port"`
	PostgresDB       string `json:"postgres_db"`
}

var (
	configuration *Config = nil
	configFile    *string = nil
)

//defined all the required flags
func init() {
	configFile = flag.String(constant.File, constant.DefaultConfig, constant.FileUsage)
}

func ResetConfiguration() {
	configuration = nil
}

func LoadAppConfiguration() {
	flag.Parse()

	if len(*configFile) == 0 {
		StopService("Mandatory arguments not provided for executing the App")
	}

	configuration = loadConfiguration(*configFile)
}

func loadConfiguration(filename string) *Config {
	configFile, err := os.Open(filename)

	if err != nil {
		StopService(err.Error())
	}

	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	e := jsonParser.Decode(&configuration)

	if e != nil {
		log.Println("Failed to parse configuration file")
		StopService(e.Error())
	}

	setDefaultConfig()

	return configuration
}

func GetAppConfiguration() *Config {
	if configuration == nil {
		log.Println("Unable to get the app configuration. Loading freshly. \t")
		LoadAppConfiguration()
	}

	log.Printf("App config ==>> %v", *configuration)

	return configuration
}

func StopService(message string) {
	p, _ := os.FindProcess(os.Getpid())
	if err := p.Signal(os.Kill); err != nil {
		log.Fatal("error killing the process while stopping the service")
	}

	log.Fatal(message)
}

func setDefaultConfig() {
	if configuration.LogLevel == "" {
		configuration.LogLevel = defaultLogLevel
	}
}
