package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/object-it/goserv/net/xhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (s *Server) HandleRecordById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getRecordByID(w, r)
	default:
		xhttp.MethodNotAllowed(w)
	}
}

func (s *Server) getRecordByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := "ID should be a number"
		log.Error(fmt.Sprintf("Server.getRecordByID - %s. %s", msg, err))
		xhttp.BadRequestWithResponse(xhttp.Response{Msg: []byte(msg), ContentType: xhttp.ContentTypeTextPlain}, w)
		return
	}

	record, err := s.RecordService.FindRecordByID(id)
	if err == sql.ErrNoRows {
		log.Error(fmt.Sprintf("Server.getRecordByID - Record with ID %d not found.", id))
		http.NotFound(w, r)
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

	xhttp.OK(xhttp.Response{Msg: bytes, ContentType: xhttp.ContentTypeApplicationJson}, w)
}
