package database

import (
	"database/sql"
	"fmt"
	"github.com/object-it/goserv/errors"
	log "github.com/sirupsen/logrus"
)

type ArtistRepository struct {
	db *sql.DB
}

// FindArtistByID does what it says
func (repository ArtistRepository) FindArtistByID(id int) (*Artist, error) {
	log.Debug("ArtistRepository.FindArtistByID - ID = ", id)

	row := repository.db.QueryRow("SELECT id, name, country FROM artists WHERE id = ?", id)
	artist := new(Artist)
	err := row.Scan(&artist.ID, &artist.Name, &artist.Country)
	if err != nil {
		return nil, errors.HandleError(log.Error, errors.New("ArtistRepository.FindArtistByID", "Error while reading data from db", err))
	}

	return artist, nil
}

// Save save an artist in the db
func (repository ArtistRepository) Save(tx *sql.Tx, artist NewArtist) (*Artist, error) {
	log.Debug("ArtistRepository.Save - ", artist.String())

	result, err := tx.Exec("INSERT INTO artists (name, country) VALUES (?, ?)", artist.Name, artist.Country)
	if err != nil {
		return nil, errors.HandleError(log.Error, errors.New("ArtistRepository.Save", fmt.Sprintf("Error while saving artist %s", artist.String()), err))
	}

	id, _ := result.LastInsertId()
	return &Artist{ID: id, Name: artist.Name, Country: artist.Country}, nil
}

// NewArtistRepository create a new ArtistRepository
func NewArtistRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{db}
}
