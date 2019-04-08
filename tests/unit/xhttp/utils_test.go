package xhttp

import (
	"github.com/gorilla/mux"
	"github.com/object-it/goserv/net/xhttp"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
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

func TestReadRequestVarIsPresent(t *testing.T) {
	var id int
	vars := map[string]string{
		"id": "1",
	}
	r := httptest.NewRequest("GET", "/artist/1", nil)
	r = mux.SetURLVars(r, vars)

	err := xhttp.ReadRequestVar(r, "id", &id)

	assert.Nil(t, err, err)
	assert.Equal(t, 1, id, "id should be equals to 1")
}

func TestReadRequestVarIsNotPresent(t *testing.T) {
	var id int
	vars := map[string]string{
		"ID": "1",
	}
	r := httptest.NewRequest("GET", "/artist/1", nil)
	r = mux.SetURLVars(r, vars)

	err := xhttp.ReadRequestVar(r, "id", &id)

	assert.NotNil(t, err, "There should have been an error")
	assert.Equal(t, "xhttp.ReadRequestVar : Not Found : id", err.Error(), "Wrong error !!!")
	assert.Equal(t, 0, id, "id should be equals to 0")
}

func TestReadRequestVarIsNotNumber(t *testing.T) {
	var id int
	vars := map[string]string{
		"id": "a",
	}
	r := httptest.NewRequest("GET", "/artist/1", nil)
	r = mux.SetURLVars(r, vars)

	err := xhttp.ReadRequestVar(r, "id", &id)

	assert.NotNil(t, err, "There should have been an error")
	assert.Contains(t, err.Error(), "xhttp.ReadRequestVar : a is not a number.", "Wrong error !!!")
	assert.Equal(t, 0, id, "id should be equals to 0")
}

func TestReadBodyToJSON_OK(t *testing.T) {
	var dummy = new(dummy)
	var r = http.Request{
		Body: testBody{"{\"a\": \"aaa\", \"b\": \"bbb\", \"c\": 1}"},
	}

	err := xhttp.ReadBodyToJSON(&r, &dummy)

	assert.Nil(t, err, err)
	assert.Equal(t, "aaa", dummy.A, "They should be equal")
	assert.Equal(t, "bbb", dummy.B, "They should be equal")
	assert.Equal(t, 1, dummy.C, "They should be equal")
}
