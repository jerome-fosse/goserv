package xhttp

import (
	"net/http"
)

type Response struct {
	Msg         []byte
	ContentType string
	Location    string
}

func BadRequest(w http.ResponseWriter) {
	w.Header().Add("Content-Type", ContentTypeTextPlain)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request !!!"))
}

func MethodNotAllowed(w http.ResponseWriter) {
	w.Header().Add("Content-Type", ContentTypeTextPlain)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method Not Allowed !!!"))
}

func BadRequestWithResponse(resp Response, w http.ResponseWriter) {
	if resp.ContentType != "" {
		w.Header().Add("Content-Type", resp.ContentType)
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp.Msg)
}

func OK(resp Response, w http.ResponseWriter) {
	if resp.ContentType != "" {
		w.Header().Add("Content-Type", resp.ContentType)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Msg)
}

func Created(resp Response, w http.ResponseWriter) {
	if resp.ContentType != "" {
		w.Header().Add("Content-Type", resp.ContentType)
	}
	if resp.Location != "" {
		w.Header().Add("Location", resp.Location)
	}
	w.WriteHeader(http.StatusCreated)
}
