package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/object-it/goserv/conf"
	log "github.com/sirupsen/logrus"
)

// OpenConnection create a new connection to the database
func OpenConnection(c conf.Configuration) *sql.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
	log.Trace("DB.OpenConnection - Trying to open a new connection to the database : " + url)

	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal("DB.OpenConnection - ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB.OpenConnection - ", err)
	}

	return db
}
