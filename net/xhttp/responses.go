package xhttp

import (
	"net/http"
)

type Response struct {
	Msg         []byte
	ContentType string
}

func BadRequest(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", ContentTypeTextPlain)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request !!!"))
}

func MethodNotAllowed(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", ContentTypeTextPlain)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method Not Allowed !!!"))
}

func BadRequestWithResponse(resp Response, w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", resp.ContentType)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp.Msg)
}

func OK(resp Response, w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", resp.ContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(resp.Msg)
}
