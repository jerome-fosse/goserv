package main

import (
	"database/sql"
	"github.com/object-it/goserv/service"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/object-it/goserv/conf"
)

type Server struct {
	config        conf.Configuration
	db            *sql.DB
	artistService *service.ArtistService
	recordService *service.RecordService
}

func (s *Server) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/artist", s.HandleArtist)
	r.HandleFunc("/artist/{id}", s.HandleArtistByID)
	r.HandleFunc("/artist/{id}/record", s.HandleArtistRecord)
	r.HandleFunc("/artist/{id}/records", s.HandleArtistRecords)

	r.HandleFunc("/record/{id}", s.HandleRecordByID)

	return r
}

func (s *Server) shutdown() {
	s.db.Close()
	log.Info("Server shutdown...")
}
