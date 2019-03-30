package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleFoo dd
func HandleFoo() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			handleGet(writer, req)
		case http.MethodPost:
			handlePost(writer, req)
		default:
			http.NotFound(writer, req)
		}
	}
}

func HandleFooById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintln(w, "Get Foo by Id -> "+vars["id"])
	}
}

func handleGet(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "Get Foo !!!!")
}

func handlePost(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(writer, "Post Foo !!!!")
}
