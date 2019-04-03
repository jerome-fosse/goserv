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
	log.Debug("ArtistRepository.FindArtistByID - ID = ", id)
	row := repository.db.QueryRow("SELECT id, name, country FROM artists WHERE id = ?", id)
	artist := new(Artist)
	err := row.Scan(&artist.ID, &artist.Name, &artist.Country)
	if err != nil {
		log.Error("ArtistRepository.FindArtistByID - ", err)
	}
	return artist, err
}

func (repository ArtistRepository) Save(artist NewArtist) (*Artist, error) {
	log.Debug("ArtistRepository.Save - ", artist.ToString())
	result, err := repository.db.Exec("INSERT INTO artists (name, country) VALUES (?, ?)", artist.Name, artist.Country)
	if err != nil {
		log.Error(fmt.Sprintf("ArtistRepository.Save - Error while saving %s. ", artist.ToString()), err)
		return nil, err
	}

	newid, _ := result.LastInsertId()
	return &Artist{ID: newid, Name: artist.Name, Country: artist.Country}, nil
}
