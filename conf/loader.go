package conf

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
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
		log.WithFields(logrus.Fields{
			"function": "loadConfigurationFile",
			"filename": filename,
		}).Fatal("Unable to load configuration file. ", err)
	}

	return parseJSONConfiguration(file)
}

func parseJSONConfiguration(file *os.File) Configuration {
	var config Configuration
	jsonParser := json.NewDecoder(file)
	jsonParser.Decode(&config)
	logrus.SetLevel(config.Logging.LogLevel())

	return config
}
