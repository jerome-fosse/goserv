package service

import (
	"database/sql"
	"github.com/object-it/goserv/database"
	log "github.com/sirupsen/logrus"
)

type RecordService struct {
	repository *database.RecordRepository
}

func NewRecordService(db *sql.DB) *RecordService {
	return &RecordService{database.NewRecordRepository(db)}
}

func (service RecordService) FindRecordByID(id int) (*database.Record, error) {
	log.Debugf("RecordService.FindRecordByID - ID = %d", id)
	return service.repository.FindRecordByID(id)
}
