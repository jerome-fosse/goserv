package conf

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var log = logrus.WithFields(logrus.Fields{
	"package": "conf",
})

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

func (l logging) LogLevel() logrus.Level {
	lvl, err := logrus.ParseLevel(l.Level)
	if err != nil {
		log.WithFields(logrus.Fields{
			"function": "LogLevel",
			"level":    l.Level,
		}).Error("Not a valid level. INFO will be used.")
		return logrus.InfoLevel
	}
	return lvl
}
