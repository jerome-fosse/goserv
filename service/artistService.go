package service

import (
	"database/sql"
	"github.com/object-it/tinyserv/database"
	log "github.com/sirupsen/logrus"
)

type ArtistService struct {
	db         *sql.DB
	repository *database.ArtistRepository
}

func NewArtistService(db *sql.DB) *ArtistService {
	return &ArtistService{db: db, repository: database.NewArtistRepository(db)}
}

func (s ArtistService) FindArtistById(id int) (*database.Artist, error) {
	return s.repository.FindArtistByID(id)
}

func (s ArtistService) SaveNewArtist(a database.NewArtist) (*database.Artist, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Error("ArtistService.SaveNewArtist - Database error : ", err)
		_ = tx.Rollback()
		return nil, err
	}

	artist, err := s.repository.Save(a)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	return artist, nil
}
