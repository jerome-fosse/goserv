package main

import (
	"github.com/gorilla/mux"
	"github.com/object-it/tinyserv/handler"
	"net/http"
)

func main() {
	srv := http.Server{
		Addr: "localhost:8080",
	}

	http.Handle("/", routes())

	srv.ListenAndServe()
}

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/foo", handler.HandleFoo())
	r.HandleFunc("/foo/{id}", handler.HandleFooById())

	return r
}
