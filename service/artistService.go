package service

import (
	"database/sql"
	"github.com/object-it/goserv/database"
	"github.com/object-it/goserv/xerror"
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
	log.Debugf("ArtistService.FindArtistById - ID = %d", id)
	return s.repository.FindArtistByID(id)
}

func (s ArtistService) SaveNewArtist(a *database.NewArtist) (int64, error) {
	log.Debugf("ArtistService.SaveNewArtist - %v", a)

	tx, err := s.db.Begin()
	if err != nil {
		return -1, xerror.HandleError(log.Error, xerror.New("ArtistService.SaveNewArtist", "Database error", err))
	}

	id, err := s.repository.Save(tx, *a)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, xerror.HandleError(log.Error, xerror.New("ArtistService.SaveNewArtist", "Database error", err))
	}

	return id, nil
}

func (s ArtistService) DeleteArtist(id int) error {
	log.Debugf("ArtistService.DeleteArtist - ID = %d", id)

	tx, err := s.db.Begin()
	if err != nil {
		return xerror.HandleError(log.Error, xerror.New("ArtistService.DeleteArtist", "Database error", err))

	}

	if err := s.repository.Delete(tx, id); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return xerror.HandleError(log.Error, xerror.New("ArtistService.DeleteArtist", "Database error", err))
	}

	return nil
}

func (s ArtistService) FindArtistDiscography(id int) (*database.Discography, error) {
	log.Debugf("ArtistService.FindArtistDiscography - ID = %d", id)
	return s.repository.FindArtistDiscography(id)
}
