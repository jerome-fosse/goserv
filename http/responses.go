package http

import (
	"net/http"
)

type Response struct {
	Msg         []byte
	ContentType string
}

func BadRequest(resp Response, w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", resp.ContentType)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp.Msg)
}

func OK(resp Response, w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", resp.ContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Msg)
}
