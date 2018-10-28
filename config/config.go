package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Config is wrapper and holder for configurations
type Config map[string]string

var logger *log.Logger

// SetLogger is setter for logger instance to be used on library
func SetLogger(li *log.Logger) {
	logger = li
}

func (conf Config) Get(key, fallback string) string {
	if _, ok := conf[key]; !ok {
		return fallback
	}

	return conf[key]
}

func ReadConfiguration() *Config {
	configuration := Config{}

	conf, err := godotenv.Read(".env")
	if err != nil {
		logger.Printf("failed to read .env file, does it exist?")

		conf, err = godotenv.Read(".env.example")
		if err != nil {
			logger.Fatalf(".env.example not found, going shutdown.")
		}

		logger.Printf("using '.env.example' as configuration file")
	}

	for key, value := range conf {
		configuration[key] = value
	}

	return &configuration
}
