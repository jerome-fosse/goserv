package main

import (
	"database/sql"
	"github.com/object-it/goserv/service"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/object-it/goserv/conf"
)

type Server struct {
	Config        conf.Configuration
	DB            *sql.DB
	ArtistService *service.ArtistService
	RecordService *service.RecordService
}

func (s *Server) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/artist", s.HandleArtist)
	r.HandleFunc("/artist/{id}", s.HandleArtistById)
	r.HandleFunc("/record/{id}", s.HandleRecordById)
	return r
}

func (s *Server) shutdown() {
	s.DB.Close()
	log.Info("Server shutdown...")
}
