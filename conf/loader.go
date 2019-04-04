package conf

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

// Load load the configuration from the config file
func Load() Configuration {
	log.Info("Server - Initializing tinyserv...")
	config := loadConfigurationFile("config.json")
	log.Debug("Server - Loading Configuration : " + config.String())
	return config
}

func loadConfigurationFile(filename string) Configuration {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("conf.loadConfigurationFile - Unable to load configuration file. ", err)
	}

	defer file.Close()

	return parseJSONConfiguration(file)
}

func parseJSONConfiguration(file *os.File) Configuration {
	var config Configuration
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	log.SetLevel(config.Logging.LogLevel())

	return config
}
