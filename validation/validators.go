package validation

import (
	"errors"
	"github.com/object-it/tinyserv/database"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

func ValidateArtistRequest(r *http.Request) (database.Artist, error) {
	var artist database.Artist

	rules := govalidator.MapData{
		"name":    []string{"required"},
		"country": []string{"required"},
	}

	messages := govalidator.MapData{
		"name":    []string{"required:name is mandatory"},
		"country": []string{"required:country is mandatory"},
	}

	opts := govalidator.Options{
		Request:  r,
		Data:     &artist,
		Rules:    rules,
		Messages: messages,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	if len(e) > 0 {
		return artist, errors.New(e.Encode())
	} else {
		return artist, nil
	}
}
