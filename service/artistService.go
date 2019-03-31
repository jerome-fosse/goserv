package service

import (
	"github.com/object-it/tinyserv/database"
)

type ArtistService struct {
	Repository *database.ArtistRepository
}

func (s ArtistService) FindArtistById(id int) (*database.Artist, error) {
	return s.Repository.FindArtistByID(id)
}
