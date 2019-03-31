package database

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type ArtistRepository struct {
	db *sql.DB
}

func NewArtistRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{db}
}

func (repository ArtistRepository) FindArtistByID(id int) (*Artist, error) {
	log.Debug(fmt.Sprintf("findArtistById : %d", id))
	row := repository.db.QueryRow("Select id, name, country from Artists where id = ?", id)
	artist := new(Artist)
	err := row.Scan(&artist.ID, &artist.Name, &artist.Country)
	return artist, err
}
