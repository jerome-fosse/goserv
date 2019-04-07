package database

import (
	"database/sql"
	"github.com/object-it/goserv/errors"

	log "github.com/sirupsen/logrus"
)

type RecordRepository struct {
	db *sql.DB
}

// NewRecordRepository créé un nouveau RecordRepository
func NewRecordRepository(db *sql.DB) *RecordRepository {
	return &RecordRepository{db}
}

func (r RecordRepository) Save(tx *sql.Tx, idArtist int, record *NewRecord) (int64, error) {
	log.Debugf("RecordRepository.Save - ID = %d, Record = %s", idArtist, record)

	r1, err := tx.Exec("INSERT INTO records(title, id_artist, year, genre, support, nb_support, label) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?)", record.Title, idArtist, record.Year, record.Genre, record.Support, record.NbSupport, record.Label)
	if err != nil {
		return -1, errors.HandleError(log.Error, errors.New("RecordRepository.Save", "database error", err))
	}

	idr, _ := r1.LastInsertId() // err toujours nil avec le driver mariadb
	for _, t := range record.Tracks {
		_, err := tx.Exec("INSERT INTO tracks (id_record, number, title, length) "+
			"VALUES(?, ?, ?, ?)", idr, t.Number, t.Title, t.Length)
		if err != nil {
			return -1, errors.HandleError(log.Error, errors.New("RecordRepository.Save", "database error", err))
		}
	}

	return r1.LastInsertId() // err toujours nil avec le driver mariadb
}

// ExistRecordByArtistIdAndTitle indique si il existe déjà un album du même titre pour le même artiste.
func (r RecordRepository) ExistRecordByArtistIdAndTitle(idArtist int, title string) (bool, error) {
	log.Debugf("RecordRepository.ExistRecordByArtistIdAndTitle - ID = %d, Title = %s", idArtist, title)

	var nb int64
	row := r.db.QueryRow(" SELECT count(*) FROM records WHERE id_artist = ? AND title = ?", idArtist, title)
	err := row.Scan(&nb)
	if err != nil {
		return true, errors.HandleError(log.Error, errors.New("RecordRepository.ExistRecordByArtistIdAndTitle", "Database error", err))
	}

	return nb > 0, nil
}

// FindRecordByID retour l'artiste dont l'id est passé en paramètre
func (r RecordRepository) FindRecordByID(id int) (*Record, error) {
	log.Debugf("RecordRepository.FindRecordByID - ID = %d", id)

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

// Delete supprimer un record et toutes ses tracks
func (r RecordRepository) Delete(tx *sql.Tx, id int) error {
	log.Debugf("RecordRepository.Delete - ID = %d", id)

	if _, err := tx.Exec("DELETE FROM records WHERE id = ?", id); err != nil {
		return errors.HandleError(log.Error, errors.New("RecordRepository.Delete", "Database error", err))
	}

	return nil
}

func (r RecordRepository) parseRowsAsRecord(rows *sql.Rows) (*Record, error) {
	record := new(Record)
	tracks := make([]Track, 0)
	var count int

	for rows.Next() {
		track := new(Track)
		if err := rows.Scan(&record.ID, &record.Title, &record.Year, &record.Genre, &record.Support, &record.NbSupport, &record.Label,
			&track.ID, &track.Number, &track.Title, &track.Length); err != nil {
			return nil, errors.HandleError(log.Error, errors.New("RecordRepository.parseRowsAsRecord", "Error while reading data from db", err))
		}
		tracks = append(tracks, *track)
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
