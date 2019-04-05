package conf

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

// Load load the configuration from the config file
func Load() Configuration {
	log.Info("Initializing goserv...")
	config := loadConfigurationFile("config.json")
	log.Debugf("Loading Configuration : %s", config.String())
	return config
}

func loadConfigurationFile(filename string) Configuration {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to load configuration file. ", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Errorf("Unexpected error : %s", err.Error())
		}
	}()

	return parseJSONConfiguration(file)
}

func parseJSONConfiguration(file *os.File) Configuration {
	var config Configuration
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	log.SetLevel(config.Logging.LogLevel())

	return config
}
