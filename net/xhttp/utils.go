package xhttp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/object-it/goserv/errors"
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
		return errors.HandleError(log.Error,
			errors.NewRootError("xhttp.ReadRequestVar", "Destination is not a pointer or is nil."))
	}

	v := mux.Vars(r)[varname]
	ddest := reflect.Indirect(rd)
	switch ddest.Kind() {
	case reflect.Int64, reflect.Int32, reflect.Int:
		i, err := strconv.Atoi(v)
		if err != nil {
			return errors.HandleError(log.Error,
				errors.New("xhttp.ReadRequestVar", v+" is not a number.", err))
		}
		ddest.SetInt(int64(i))
	case reflect.String:
		ddest.SetString(v)
	default:
		return errors.HandleError(log.Error,
			errors.NewRootError("xhttp.ReadRequestVar", "Could not read variable. Format unknown."))
	}

	return nil
}

func ReadBodyToJSON(r *http.Request, dest interface{}) error {
	rd := reflect.ValueOf(dest)
	if rd.Kind() != reflect.Ptr || rd.IsNil() {
		return errors.HandleError(log.Error,
			errors.NewRootError("xhttp.ReadBodyToJSON", "Destination is not a pointer or is nil."))
	}

	buffer, err := ioutil.ReadAll(r.Body)
	log.Debugf("xhttp.ReadBodyToJSON - Body is : %s", buffer)
	if err != nil && err != io.EOF {
		return errors.HandleError(log.Error,
			errors.New("xhttp.ReadBodyToJSON", "unable to read JSON", err))
	}

	if err := json.Unmarshal(buffer, dest); err != nil {
		log.Error("RecordHandler.postRecord - Unable to parse Json message. ", err)
		return errors.HandleError(log.Error,
			errors.New("xhttp.ReadBodyToJSON", "unable to parse JSON", err))
	}

	return nil
}
