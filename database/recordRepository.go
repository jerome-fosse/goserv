package database

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type RecordRepository struct {
	db *sql.DB
}

func NewRecordRepository(db *sql.DB) *RecordRepository {
	return &RecordRepository{db}
}

func (r *RecordRepository) FindRecordByID(id int) (*Record, error) {
	log.Info("RecordRepository.FindRecordByID - ID = ", id)
	rows, err := r.db.Query(
		"SELECT rec.id, rec.title, rec.year, rec.genre, rec.support, rec.nb_support, rec.label, "+
			"tra.id as id_track, tra.number, tra.title, tra.length "+
			"FROM records rec LEFT JOIN tracks tra on rec.id = tra.id_record "+
			"WHERE rec.id = ? "+
			"ORDER BY tra.number ASC", id)
	if err != nil {
		log.Error("RecordRepository.FindRecordByID - ", err)
		return nil, err
	}
	defer rows.Close()

	return r.parseRowsAsRecord(rows)
}

func (r *RecordRepository) parseRowsAsRecord(rows *sql.Rows) (*Record, error) {
	record := new(Record)
	tracks := make([]*Track, 0)

	for rows.Next() {
		track := new(Track)
		if err := rows.Scan(&record.ID, &record.Title, &record.Year, &record.Genre, &record.Support, &record.NbSupport, &record.Label,
			&track.ID, &track.Number, &track.Title, &track.Length); err != nil {
			log.Error("RecordRepository.parseRowsAsRecord - ", err)
			return nil, err
		}
		tracks = append(tracks, track)
	}

	record.Tracks = tracks
	return record, nil
}
