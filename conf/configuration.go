package conf

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Configuration is a structure containing the configuuration of the application
type Configuration struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string `json:"dbname"`
	}
	Logging logging
}

type logging struct {
	Level string `json:"level"`
}

// ToString create a string representation of the configuration
func (c Configuration) ToString() string {
	var conflines [2]string

	conflines[0] = fmt.Sprintf("Database: [host: %s - port: %d - dbname: %s - user: %s - passord: %s]",
		c.Database.Host, c.Database.Port, c.Database.Name, c.Database.User, c.Database.Password)
	conflines[1] = fmt.Sprintf("Logging : [Level: %s]", c.Logging.Level)

	return fmt.Sprint(conflines)
}

func (l logging) LogLevel() log.Level {
	lvl, err := log.ParseLevel(l.Level)
	if err != nil {
		log.Error(fmt.Sprintf("%s is not a valid log level. INFO will be used.", l.Level))
		return log.InfoLevel
	}
	return lvl
}
