package main

import (
	"database/sql"
	"encoding/json"
	"github.com/object-it/goserv/errors"
	"github.com/object-it/goserv/net/xhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// HandleRecordByID handler qui gère les requetes sur la ressource /record/{id}
func (s *Server) HandleRecordByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getRecordByID(w, r)
	default:
		xhttp.MethodNotAllowed(w)
	}
}

// GET /record/{id}
func (s *Server) getRecordByID(w http.ResponseWriter, r *http.Request) {
	var id int

	if err := xhttp.ReadRequestVar(r, "id", &id); err != nil {
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	log.Infof("HandleRecord.getRecordByID - ID = %d", id)

	record, err := s.RecordService.FindRecordByID(id)
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

	bytes, err := json.Marshal(record)
	if err != nil {
		log.Error("Server.getRecordByID - Unexpected error. ", err)
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(err.Error()), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	xhttp.OK(xhttp.Response{Msg: bytes, ContentType: xhttp.ContentTypeApplicationJson}, w)
}
