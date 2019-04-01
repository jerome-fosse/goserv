package service

import (
	"database/sql"

	"github.com/object-it/tinyserv/database"
)

type ArtistService struct {
	repository *database.ArtistRepository
}

func NewArtistService(db *sql.DB) *ArtistService {
	return &ArtistService{database.NewArtistRepository(db)}
}

func (s ArtistService) FindArtistById(id int) (*database.Artist, error) {
	return s.repository.FindArtistByID(id)
}
