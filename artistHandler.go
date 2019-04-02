package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/object-it/tinyserv/net/xhttp"
	"github.com/object-it/tinyserv/validation"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (s *Server) HandleArtist(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.postArtist(w, r)
	default:
		xhttp.MethodNotAllowed(w, r)
	}
}

func (s *Server) HandleArtistById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getArtistByID(w, r)
	default:
		xhttp.MethodNotAllowed(w, r)
	}
}

func (s *Server) postArtist(w http.ResponseWriter, r *http.Request) {
	log.Info("ArtistHandler.postArtist")

	artist, err := validation.ValidateArtistRequest(r)
	if err != nil {
		log.Error("ArtistHandler.postArtist - Payload invalid. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain,}, w, r)
		return
	}

	fmt.Printf("Artist : %s - %s", artist.Name, artist.Country)
}

func (s *Server) getArtistByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "ID should be a number"
		log.Error(fmt.Sprintf("Server.getArtistById - %s. %s", msg, err))
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(msg), ContentType: xhttp.ContentTypeTextPlain}, w, r)
		return
	}

	a, err := s.ArtistService.FindArtistById(id)
	if err == sql.ErrNoRows {
		log.Error(fmt.Sprintf("Server.getArtistById - Artist with ID %d not found.", id))
		http.NotFound(w, r)
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

	xhttp.OK(xhttp.Response{Msg: bytes, ContentType: xhttp.ContentTypeApplicationJson}, w, r)
}

