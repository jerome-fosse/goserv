package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/object-it/tinyserv/service"

	h "github.com/object-it/tinyserv/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/object-it/tinyserv/conf"
)

type Server struct {
	Config        conf.Configuration
	DB            *sql.DB
	ArtistService *service.ArtistService
}

func (s *Server) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/artist/{id}", s.HandleArtistById)
	return r
}

func (s *Server) shutdown() {
	s.DB.Close()
	log.Info("Server shutdown...")
}

func (s *Server) HandleArtistById(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		s.getArtistByID(w, req)
	default:
		http.NotFound(w, req)
	}
}

func (s *Server) getArtistByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "ID should be a number."
		log.Error(fmt.Sprintf("Server.getArtistById -  %s. %s", msg, err))
		h.BadRequest(h.Response{Msg: []byte(msg), ContentType: "test/plain"}, w, req)
		return
	}

	a, err := s.ArtistService.FindArtistById(id)
	if err == sql.ErrNoRows {
		log.Error(fmt.Sprintf("Server.getArtistById - Artist with ID %d not found.", id))
		http.NotFound(w, req)
		return
	} else if err != nil {
		log.Error("Server.getArtistById - Unexpected error.", err)
		http.Error(w, err.Error(), 400)
		return
	}

	bytes, err := json.Marshal(a)
	if err != nil {
		log.Error("Server.getArtistById - Unexpected error. ", err)
		http.Error(w, err.Error(), 400)
		return
	}

	h.OK(h.Response{Msg: bytes, ContentType: "application/json"}, w, req)
}
