package main

import (
	"net/http"

	"github.com/object-it/tinyserv/service"

	"github.com/object-it/tinyserv/conf"
	"github.com/object-it/tinyserv/database"
	log "github.com/sirupsen/logrus"
)

var server Server

func init() {
	log.Info("Server - Initializing tinyserv...")
	c := conf.Load()
	log.Debug("Server - Loading Configuration : " + c.ToString())

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
		Addr: "localhost:8080",
	}

	http.Handle("/", server.routes())
	srv.RegisterOnShutdown(server.shutdown)
	srv.ListenAndServe()
}
