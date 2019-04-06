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

type NewArtist struct {
	Name    string `json:"name" valid:"required~Name is mandatory"`
	Country string `json:"country" valid:"required~Country is mandatory"`
}

func (a NewArtist) String() string {
	return fmt.Sprintf("Artist: [Name = %s - Country = %s]", a.Name, a.Country)
}

type NewRecord struct {
	Title     string     `json:"title" valid:"required~The record's title is mandatory"`
	Year      int64      `json:"year"`
	Genre     string     `json:"genre"`
	Support   string     `json:"support"`
	NbSupport int64      `json:"nb_supports"`
	Label     string     `json:"label"`
	Tracks    []NewTrack `json:"tracks" valid:"required~A record should have at least one track"`
}

func (r NewRecord) String() string {
	return fmt.Sprintf("Record: [Title: %s - Year: %d - Genre: %s - Support: %s - NbSupport: %d - Label: %s - Tracks: %v]",
		r.Title, r.Year, r.Genre, r.Support, r.NbSupport, r.Label, r.Tracks)
}

type NewTrack struct {
	Number int64  `json:"number" valid:"required~The tracks's number is mandatory"`
	Title  string `json:"title" valid:"required~The track's title is mandatory"`
	Length int64  `json:"length"`
}

func (t NewTrack) String() string {
	return fmt.Sprintf("Track : [Number: %d - Title: %s - Length: %d]", t.Number, t.Title, t.Length)
}
