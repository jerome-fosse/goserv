package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/object-it/goserv/database"
	"github.com/object-it/goserv/errors"
	"github.com/object-it/goserv/net/xhttp"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

// HandleArtist request to path /artist
func (s *Server) HandleArtist(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.postArtist(w, r)
	default:
		xhttp.MethodNotAllowed(w)
	}
}

// HandleArtistByID request to path /artist/{id}
func (s *Server) HandleArtistByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getArtistByID(w, r)
	default:
		xhttp.MethodNotAllowed(w)
	}
}

// HandleArtistDiscography request to path /artist/{id}/discography
func (s *Server) HandleArtistDiscography(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getArtistDiscography(w, r)
	default:
		xhttp.MethodNotAllowed(w)
	}
}

func (s *Server) postArtist(w http.ResponseWriter, r *http.Request) {
	log.Info("ArtistHandler.postArtist")

	artist := new(database.NewArtist)

	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		log.Error("ArtistHandler.postArtist - Unable to read Json message. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	if err := json.Unmarshal(buffer, &artist); err != nil {
		log.Error("ArtistHandler.postArtist - Unable to parse Json message. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	if _, err := govalidator.ValidateStruct(artist); err != nil {
		log.Error("ArtistHandler.postArtist - Unable to validate Json message. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	a, err := s.ArtistService.SaveNewArtist(artist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Info("ArtistHandler.postArtist - saved : " + a.String())
	xhttp.Created(xhttp.Response{Location: "/artist/" + strconv.FormatInt(a.ID, 10)}, w)
}

func (s *Server) getArtistByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "ID should be a number"
		log.Error(fmt.Sprintf("ArtistHandler.getArtistById - %s. %s", msg, err))
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(msg), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	a, err := s.ArtistService.FindArtistById(id)
	if err != nil {
		switch errors.Cause(err) {
		case sql.ErrNoRows:
			http.NotFound(w, r)
			return
		default:
			xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
			return
		}
	}

	bytes, err := json.Marshal(a)
	if err != nil {
		log.Error("ArtistHandler.getArtistById - Unexpected error. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	xhttp.OK(xhttp.Response{Msg: bytes, ContentType: xhttp.ContentTypeApplicationJson}, w)
}

func (s *Server) getArtistDiscography(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "ID should be a number"
		log.Error(fmt.Sprintf("ArtistHandler.getArtistDiscography - %s. %s", msg, err))
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(msg), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	d, err := s.ArtistService.FindArtistDiscography(id)
	if err != nil {
		switch errors.Cause(err) {
		case sql.ErrNoRows:
			http.NotFound(w, r)
			return
		default:
			xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
			return
		}
	}

	bytes, err := json.Marshal(d)
	if err != nil {
		log.Error("ArtistHandler.getArtistDiscography - Unexpected error. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	xhttp.OK(xhttp.Response{Msg: bytes, ContentType: xhttp.ContentTypeApplicationJson}, w)
}
