package database

import (
	"fmt"
)

type Artist struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Country NullString `json:"country"`
}

func (a Artist) String() string {
	return fmt.Sprintf("Artist: [ID = %d - Name = %s - Country = %s]", a.ID, a.Name, a.Country.String)
}

type NewArtist struct {
	Name    string `json:"name" valid:"required~name is mandatory"`
	Country string `json:"country" valid:"required~country is mandatory"`
}

func (a NewArtist) String() string {
	return fmt.Sprintf("Artist: [Name = %s - Country = %s]", a.Name, a.Country)
}

type Record struct {
	ID        int64      `json:"id"`
	Title     string     `json:"title"`
	Year      NullInt64  `json:"year"`
	Genre     NullString `json:"genre"`
	Support   NullString `json:"support"`
	NbSupport NullInt64  `json:"nb_supports"`
	Label     NullString `json:"label"`
	Tracks    []Track    `json:"tracks"`
}

type Track struct {
	ID     int64     `json:"id"`
	Number int64     `json:"number"`
	Title  string    `json:"title"`
	Length NullInt64 `json:"length"`
}

type Discography struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Country NullString `json:"country"`
	Records []Record   `json:"records"`
}
