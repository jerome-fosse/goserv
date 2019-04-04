package main

import (
	"net/http"

	"github.com/object-it/goserv/service"

	"github.com/object-it/goserv/conf"
	"github.com/object-it/goserv/database"
	log "github.com/sirupsen/logrus"
)

var server Server

func init() {
	c := conf.Load()
	db := database.OpenConnection(c)

	server = Server{
		Config:        c,
		DB:            db,
		ArtistService: service.NewArtistService(db),
		RecordService: service.NewRecordService(db),
	}
}

func main() {
	log.Info("Server - Starting tinyserv...")

	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: server.routes(),
	}
	srv.RegisterOnShutdown(server.shutdown)

	log.Fatal(srv.ListenAndServe())
}
