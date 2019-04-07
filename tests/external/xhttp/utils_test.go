package xhttp

import (
	"context"
	"github.com/object-it/goserv/net/xhttp"
	"io"
	"net/http"
	"testing"
)

type dummy struct {
	A string `json:"a"`
	B string `json:"b"`
	C int    `json:"c"`
}

type testBody struct {
	body string
}

func (b testBody) Read(p []byte) (n int, err error) {
	copy(p, []byte(b.body))
	return len(b.body), io.EOF
}

func (b testBody) Close() error {
	return nil
}

func testReadRequestVarIsPresent(t *testing.T) {
	r := http.Request{Method: "GET", RequestURI: "/artist/1"}
	ctx := context.WithValue(r.Context(), "id", "1")
	ctx2 := context.WithValue(ctx, "", "")
	var id int

	err := xhttp.ReadRequestVar(r.WithContext(ctx2), "id", &id)
	if err != nil {
		t.Error("Unable to read var id in the http.Request")
	}

	if id != 1 {
		t.Errorf("id should have bean equals to 1 and is equal to %d", id)
	}
}

func TestReadBodyToJSON_OK(t *testing.T) {
	var dummy = new(dummy)
	var r = http.Request{
		Body: testBody{"{\"a\": \"aaa\", \"b\": \"bbb\", \"c\": 1}"},
	}

	if err := xhttp.ReadBodyToJSON(&r, &dummy); err != nil {
		t.Error(err.Error())
	}

	if dummy.A != "aaa" {
		t.Errorf("dummy.a : expected [aaa]   current [%s]", dummy.A)
	}

	if dummy.B != "bbb" {
		t.Errorf("dummy.b : expected [bbb]   current [%s]", dummy.B)
	}

	if dummy.C != 1 {
		t.Errorf("dummy.c : expected [1]   current [%d]", dummy.C)
	}
}
