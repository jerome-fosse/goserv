package service

import (
	"database/sql"

	"github.com/object-it/goserv/database"
)

type RecordService struct {
	repository *database.RecordRepository
}

func NewRecordService(db *sql.DB) *RecordService {
	return &RecordService{database.NewRecordRepository(db)}
}

func (service RecordService) FindRecordByID(id int) (*database.Record, error) {
	return service.repository.FindRecordByID(id)
}
