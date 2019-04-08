package service

import (
	"database/sql"
	"fmt"
	"github.com/object-it/goserv/database"
	"github.com/object-it/goserv/xerrors"
	log "github.com/sirupsen/logrus"
)

type RecordService struct {
	db         *sql.DB
	artistRepo *database.ArtistRepository
	recordRepo *database.RecordRepository
}

func NewRecordService(db *sql.DB) *RecordService {
	return &RecordService{db: db, artistRepo: database.NewArtistRepository(db),
		recordRepo: database.NewRecordRepository(db)}
}

func (s RecordService) FindRecordByID(id int) (*database.Record, error) {
	log.Debugf("RecordService.FindRecordByID - ID = %d", id)
	return s.recordRepo.FindRecordByID(id)
}

func (s RecordService) DeleteRecord(id int) error {
	log.Debugf("RecordService.DeleteRecord - ID = %d", id)

	tx, err := s.db.Begin()
	if err != nil {
		return xerrors.HandleError(log.Error, xerrors.New("RecordService.DeleteRecord", "Database error", err))
	}

	if err := s.recordRepo.Delete(tx, id); err != nil {
		_ = tx.Rollback()
		return xerrors.HandleError(log.Error, xerrors.New("RecordService.DeleteRecord", "Database error", err))
	}

	if err := tx.Commit(); err != nil {
		return xerrors.HandleError(log.Error, xerrors.New("RecordService.DeleteRecord", "Database error", err))
	}

	return nil
}

func (s RecordService) SaveRecordForArtist(idArtist int, record *database.NewRecord) (int64, error) {
	log.Debugf("RecordService.SaveRecordForArtist - ID = %d, Record = %s", idArtist, record)

	_, err := s.artistRepo.FindArtistByID(idArtist)
	if err != nil {
		return -1, err
	}

	exist, err := s.recordRepo.ExistRecordByArtistIdAndTitle(idArtist, record.Title)
	if err != nil {
		return -1, err
	}

	if exist {
		return -1, xerrors.HandleError(log.Error, xerrors.NewRootError("RecordService.SaveRecordForArtist",
			fmt.Sprintf("The record %s already exist for artit with id %d", record.Title, idArtist)))
	}

	tx, err := s.db.Begin()
	if err != nil {
		return -1, xerrors.HandleError(log.Error, xerrors.New("RecordService.SaveRecordForArtist", "Database error", err))
	}

	idr, err := s.recordRepo.Save(tx, idArtist, record)
	if err != nil {
		_ = tx.Rollback()
		return -1, err
	}

	if err := tx.Commit(); err != nil {
		return -1, xerrors.HandleError(log.Error, xerrors.New("RecordService.SaveRecordForArtist", "Database error", err))
	}

	return idr, nil
}
