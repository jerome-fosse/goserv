package main

import (
	"github.com/gorilla/mux"
	"github.com/object-it/tinyserv/configuration"
	"github.com/object-it/tinyserv/handler"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var c configuration.Configuration

func init() {
	c := configuration.Load()
	log.SetLevel(c.Logging.LogLevel())
	log.Debug("Loading Configuration : " + c.ToString())
}

func main() {
	log.Info("Starting tinyserv...")

	srv := http.Server{
		Addr: "localhost:8080",
	}

	http.Handle("/", routes())

	srv.ListenAndServe()
}

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/foo", handler.HandleFoo())
	r.HandleFunc("/foo/{id}", handler.HandleFooById())

	return r
}
