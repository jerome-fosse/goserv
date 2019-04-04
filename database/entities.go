package database

import "fmt"

type Artist struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func (a Artist) String() string {
	return fmt.Sprintf("Artist: [ID = %d - Name = %s - Country = %s]", a.ID, a.Name, a.Country)
}

type NewArtist struct {
	Name    string `json:"name" valid:"required~name is mandatory"`
	Country string `json:"country" valid:"required~country is mandatory"`
}

func (a NewArtist) String() string {
	return fmt.Sprintf("Artist: [Name = %s - Country = %s]", a.Name, a.Country)
}

type Record struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Year      *int     `json:"year"`
	Genre     *string  `json:"genre"`
	Support   *string  `json:"support"`
	NbSupport *int     `json:"nb_supports"`
	Label     *string  `json:"label"`
	Tracks    []*Track `json:"tracks"`
}

type Track struct {
	ID     int    `json:"id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Length *int   `json:"length"`
}
