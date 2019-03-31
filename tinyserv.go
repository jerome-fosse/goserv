package main

import (
	"github.com/gorilla/mux"
	"github.com/object-it/tinyserv/conf"
	"github.com/object-it/tinyserv/database"
	"github.com/object-it/tinyserv/handler"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var server Server

func init() {
	log.Info("Initializing tinyserv...")
	c := conf.Load()
	log.Debug("Loading Configuration : " + c.ToString())

	db := database.OpenConnection(c)

	server = Server{
		Config: c,
		DB:     db,
	}
}

func main() {
	log.Info("Starting tinyserv...")

	srv := http.Server{
		Addr: "localhost:8080",
	}

	http.Handle("/", server.routes())

	srv.ListenAndServe()
}

func (s *Server) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/foo", handler.HandleFoo())
	r.HandleFunc("/foo/{id}", handler.HandleFooById())

	r.HandleFunc("/artist/{id}", s.HandleArtistById)
	return r
}
