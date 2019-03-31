package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/object-it/tinyserv/conf"
	log "github.com/sirupsen/logrus"
)

// OpenConnection create a new connection to the database
func OpenConnection(c conf.Configuration) *sql.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
	log.Trace("Trying to open a new connection to the database : " + url)

	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err)
	}

	err2 := db.Ping()
	if err2 != nil {
		log.Fatal(err2)
	}

	return db
}
