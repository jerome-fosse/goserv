package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/object-it/goserv/conf"
	"github.com/object-it/goserv/database"
	"github.com/object-it/goserv/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var server Server

func init() {
	c := conf.Load()
	db := database.OpenConnection(c)

	server = Server{
		config:        c,
		db:            db,
		artistService: service.NewArtistService(db),
		recordService: service.NewRecordService(db),
	}
}

func main() {
	log.Info("Server - Starting goserv...")

	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: server.routes(),
	}
	srv.RegisterOnShutdown(server.shutdown)

	log.Fatal(srv.ListenAndServe())
}
