package xhttp

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Response struct {
	Msg         []byte
	ContentType string
	Location    string
}

// BadRequest
func BadRequest(w http.ResponseWriter) {
	w.Header().Add("Content-Type", ContentTypeTextPlain)
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write([]byte("Bad Request !!!")); err != nil {
		log.Error(err)
	}
}

// BadRequestWithResponse
func BadRequestWithResponse(resp Response, w http.ResponseWriter) {
	if resp.ContentType != "" {
		w.Header().Add("Content-Type", resp.ContentType)
	}
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write(resp.Msg); err != nil {
		log.Error(err)
	}
}

// MethodNotAllowed
func MethodNotAllowed(w http.ResponseWriter) {
	w.Header().Add("Content-Type", ContentTypeTextPlain)
	w.WriteHeader(http.StatusMethodNotAllowed)
	if _, err := w.Write([]byte("Method Not Allowed !!!")); err != nil {
		log.Error(err)
	}
}

// OK
func OK(resp Response, w http.ResponseWriter) {
	if resp.ContentType != "" {
		w.Header().Add("Content-Type", resp.ContentType)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp.Msg); err != nil {
		log.Error(err)
	}
}

// Created
func Created(resp Response, w http.ResponseWriter) {
	if resp.ContentType != "" {
		w.Header().Add("Content-Type", resp.ContentType)
	}
	if resp.Location != "" {
		w.Header().Add("Location", resp.Location)
	}
	w.WriteHeader(http.StatusCreated)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)

}
