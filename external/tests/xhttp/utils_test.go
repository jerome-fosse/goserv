package xhttp

import (
	"context"
	"github.com/object-it/goserv/net/xhttp"
	"net/http"
	"testing"
)

func testReadRequestVarIsPresent(t *testing.T) {
	r := http.Request{Method: "GET", RequestURI: "/artist/1"}
	ctx := context.WithValue(r.Context(), "id", "1")
	var id int

	err := xhttp.ReadRequestVar(r.WithContext(ctx), "id", &id)
	if err != nil {
		t.Error("Unable to read var id in the http.Request")
	}

	if id != 1 {
		t.Errorf("id should have bean equals to 1 and is equal to %d", id)
	}
}
