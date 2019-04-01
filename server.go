package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/object-it/tinyserv/service"

	"github.com/object-it/tinyserv/net/xhttp"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/object-it/tinyserv/conf"
)

type Server struct {
	Config        conf.Configuration
	DB            *sql.DB
	ArtistService *service.ArtistService
	RecordService *service.RecordService
}

func (s *Server) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/artist/{id}", s.HandleArtistById)
	r.HandleFunc("/record/{id}", s.HandleRecordById)
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
		xhttp.MethodNotAllowed(w, req)
	}
}

func (s *Server) HandleRecordById(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		s.getRecordByID(w, req)
	default:
		xhttp.MethodNotAllowed(w, req)
	}
}

func (s *Server) getArtistByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "ID should be a number"
		log.Error(fmt.Sprintf("Server.getArtistById - %s. %s", msg, err))
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(msg), ContentType: xhttp.ContentTypeTextPlain}, w, req)
		return
	}

	a, err := s.ArtistService.FindArtistById(id)
	if err == sql.ErrNoRows {
		log.Error(fmt.Sprintf("Server.getArtistById - Artist with ID %d not found.", id))
		http.NotFound(w, req)
		return
	} else if err != nil {
		log.Error("Server.getArtistById - Unexpected error.", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(a)
	if err != nil {
		log.Error("Server.getArtistById - Unexpected error. ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	xhttp.OK(xhttp.Response{Msg: bytes, ContentType: xhttp.ContentTypeApplicationJson}, w, req)
}

func (s *Server) getRecordByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "ID should be a number"
		log.Error(fmt.Sprintf("Server.getRecordByID - %s. %s", msg, err))
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(msg), ContentType: xhttp.ContentTypeTextPlain}, w, req)
		return
	}

	record, err := s.RecordService.FindRecordByID(id)
	if err == sql.ErrNoRows {
		log.Error(fmt.Sprintf("Server.getRecordByID - Record with ID %d not found.", id))
		http.NotFound(w, req)
		return
	} else if err != nil {
		log.Error("Server.getRecordByID - Unexpected error.", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(record)
	if err != nil {
		log.Error("Server.getRecordByID - Unexpected error. ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	xhttp.OK(xhttp.Response{Msg: bytes, ContentType: xhttp.ContentTypeApplicationJson}, w, req)
}
