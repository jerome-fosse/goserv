package xhttp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/object-it/goserv/xerrors"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

func ReadRequestVar(r *http.Request, varname string, dest interface{}) error {
	rd := reflect.ValueOf(dest)
	if rd.Kind() != reflect.Ptr || rd.IsNil() {
		return xerrors.HandleError(log.Error,
			xerrors.NewRootError("xhttp.ReadRequestVar", "Destination is not a pointer or is nil."))
	}

	v := mux.Vars(r)[varname]
	if v == "" {
		return xerrors.HandleError(log.Error,
			xerrors.NewRootError("xhttp.ReadRequestVar", "Not Found : "+varname))
	}

	ddest := reflect.Indirect(rd)
	switch ddest.Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int:
		i, err := strconv.Atoi(v)
		if err != nil {
			return xerrors.HandleError(log.Error,
				xerrors.New("xhttp.ReadRequestVar", v+" is not a number.", err))
		}
		ddest.SetInt(int64(i))
	case reflect.String:
		ddest.SetString(v)
	default:
		return xerrors.HandleError(log.Error,
			xerrors.NewRootError("xhttp.ReadRequestVar", "Could not read variable. Format unknown."))
	}

	return nil
}

func ReadBodyToJSON(r *http.Request, dest interface{}) error {
	rd := reflect.ValueOf(dest)
	if rd.Kind() != reflect.Ptr || rd.IsNil() {
		return xerrors.HandleError(log.Error,
			xerrors.NewRootError("xhttp.ReadBodyToJSON", "Destination is not a pointer or is nil."))
	}

	buffer, err := ioutil.ReadAll(r.Body)
	log.Debugf("xhttp.ReadBodyToJSON - Body is : %s", buffer)
	if err != nil && err != io.EOF {
		return xerrors.HandleError(log.Error,
			xerrors.New("xhttp.ReadBodyToJSON", "unable to read JSON", err))
	}

	if err := json.Unmarshal(buffer, dest); err != nil {
		return xerrors.HandleError(log.Error,
			xerrors.New("xhttp.ReadBodyToJSON", "unable to parse JSON", err))
	}

	return nil
}
