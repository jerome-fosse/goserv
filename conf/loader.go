package conf

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

// Load load the configuration from the config file
func Load() Configuration {
	config := loadConfigurationFile("config.json")
	return config
}

func loadConfigurationFile(filename string) Configuration {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal("Unable to find " + filename + " file.")
	}

	return parseJSONConfiguration(file)
}

func parseJSONConfiguration(file *os.File) Configuration {
	var config Configuration
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	log.SetLevel(config.Logging.LogLevel())

	return config
}
