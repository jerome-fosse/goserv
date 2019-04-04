package database

import (
	"database/sql"
	"github.com/object-it/goserv/errors"

	log "github.com/sirupsen/logrus"
)

type RecordRepository struct {
	db *sql.DB
}

// NewRecordRepository create a new RecordRepository
func NewRecordRepository(db *sql.DB) *RecordRepository {
	return &RecordRepository{db}
}

// FindRecordByID does what it says
func (r *RecordRepository) FindRecordByID(id int) (*Record, error) {
	log.Info("RecordRepository.FindRecordByID - ID = ", id)
	rows, err := r.db.Query(
		"SELECT rec.id, rec.title, rec.year, rec.genre, rec.support, rec.nb_support, rec.label, "+
			"tra.id as id_track, tra.number, tra.title, tra.length "+
			"FROM records rec LEFT JOIN tracks tra on rec.id = tra.id_record "+
			"WHERE rec.id = ? "+
			"ORDER BY tra.number ASC", id)
	if err != nil {
		return nil, errors.HandleError(log.Error, errors.New("RecordRepository.FindRecordByID", "Database error", err))
	}
	defer rows.Close()

	return r.parseRowsAsRecord(rows)
}

func (r *RecordRepository) parseRowsAsRecord(rows *sql.Rows) (*Record, error) {
	record := new(Record)
	tracks := make([]*Track, 0)
	var count int

	for rows.Next() {
		track := new(Track)
		if err := rows.Scan(&record.ID, &record.Title, &record.Year, &record.Genre, &record.Support, &record.NbSupport, &record.Label,
			&track.ID, &track.Number, &track.Title, &track.Length); err != nil {
			return nil, errors.HandleError(log.Error, errors.New("RecordRepository.parseRowsAsRecord", "Error while reading data from db", err))
		}
		tracks = append(tracks, track)
		count++
	}

	if count == 0 {
		return nil, errors.HandleError(log.Error, errors.New("RecordRepository.parseRowsAsRecord", "Error while reading data from db", sql.ErrNoRows))
	}

	err := rows.Err()
	if err != nil {
		return nil, errors.HandleError(log.Error, errors.New("RecordRepository.parseRowsAsRecord", "Error while reading data from db", err))
	}

	record.Tracks = tracks
	return record, nil
}
