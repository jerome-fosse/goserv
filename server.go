package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/object-it/tinyserv/conf"
	"github.com/object-it/tinyserv/database"
	"github.com/object-it/tinyserv/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Server struct {
	Config conf.Configuration
	DB     *sql.DB
}

func (s *Server) HandleArtistById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "ID Artist should be a number"
		log.Error(msg, err)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
	} else {
		serv := s.newArtistService()
		a, _ := serv.FindArtistById(id)
		bytes, _ := json.Marshal(a)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func (s *Server) newArtistService() *service.ArtistService {
	return &service.ArtistService{database.NewArtistRepository(s.DB)}
}
