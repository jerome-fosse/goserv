package main

import (
	"database/sql"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/object-it/goserv/database"
	"github.com/object-it/goserv/errors"
	"github.com/object-it/goserv/net/xhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// HandleArtist gère les requetes sur l'url /artist
func (s *Server) HandleArtist(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.postArtist(w, r)
	default:
		xhttp.MethodNotAllowed(w)
	}
}

// HandleArtistByID gère les requetes sur l'url /artist/{id}
func (s *Server) HandleArtistByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getArtistByID(w, r)
	default:
		xhttp.MethodNotAllowed(w)
	}
}

// HandleArtistRecords gère les requetes sur l'url /artist/{id}/records
func (s *Server) HandleArtistRecords(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getArtistRecords(w, r)
	default:
		xhttp.MethodNotAllowed(w)
	}
}

// HandleArtistRecords gère les requetes sur l'url /artist/{id}/record
func (s *Server) HandleArtistRecord(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.postArtistRecord(w, r)
	default:
		xhttp.MethodNotAllowed(w)
	}
}

// POST /artist
func (s *Server) postArtist(w http.ResponseWriter, r *http.Request) {
	log.Info("ArtistHandler.postArtist")

	artist := new(database.NewArtist)

	if err := xhttp.ReadBodyToJSON(r, artist); err != nil {
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	if _, err := govalidator.ValidateStruct(artist); err != nil {
		log.Error("ArtistHandler.postArtist - Unable to validate Json message. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	id, err := s.artistService.SaveNewArtist(artist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Infof("ArtistHandler.postArtist - saved : Artist = %v, ID = %d", artist, id)
	xhttp.Created(xhttp.Response{Location: "/artist/" + strconv.FormatInt(id, 10)}, w)
}

// GET /artist/{id}
func (s *Server) getArtistByID(w http.ResponseWriter, r *http.Request) {
	var id int

	if err := xhttp.ReadRequestVar(r, "id", &id); err != nil {
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	log.Infof("ArtistHandler.getArtistById - ID = %d", id)

	a, err := s.artistService.FindArtistById(id)
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

// GET /artist/{id}/records
func (s *Server) getArtistRecords(w http.ResponseWriter, r *http.Request) {

	var id int

	if err := xhttp.ReadRequestVar(r, "id", &id); err != nil {
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	log.Infof("ArtistHandler.getArtistRecords - Artist ID = %d", id)

	d, err := s.artistService.FindArtistDiscography(id)
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
		log.Error("ArtistHandler.getArtistRecords - Unexpected error. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	xhttp.OK(xhttp.Response{Msg: bytes, ContentType: xhttp.ContentTypeApplicationJson}, w)
}

// POST /artist/{1}/record
func (s *Server) postArtistRecord(w http.ResponseWriter, r *http.Request) {
	var (
		id     int
		record = new(database.NewRecord)
	)

	if err := xhttp.ReadRequestVar(r, "id", &id); err != nil {
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	log.Infof("ArtistHandler.postArtistRecord - Artist ID = %d", id)

	if err := xhttp.ReadBodyToJSON(r, record); err != nil {
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	if _, err := govalidator.ValidateStruct(record); err != nil {
		log.Error("RecordHandler.postArtistRecord - Unable to validate Json message. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	idr, err := s.recordService.SaveRecordForArtist(id, record)
	if err != nil {
		log.Error("RecordHandler.postArtistRecord - Unable to save record into db. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	xhttp.Created(xhttp.Response{Location: "/record/" + strconv.FormatInt(idr, 10), ContentType: xhttp.ContentTypeApplicationJson}, w)
}
