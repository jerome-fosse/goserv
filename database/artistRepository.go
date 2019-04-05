package database

import (
	"database/sql"
	"fmt"
	"github.com/object-it/goserv/errors"
	log "github.com/sirupsen/logrus"
)

type ArtistRepository struct {
	db *sql.DB
}

// FindArtistByID does what it says
func (r ArtistRepository) FindArtistByID(id int) (*Artist, error) {
	log.Debug("ArtistRepository.FindArtistByID - ID = ", id)

	row := r.db.QueryRow("SELECT id, name, country FROM artists WHERE id = ?", id)
	artist := new(Artist)
	err := row.Scan(&artist.ID, &artist.Name, &artist.Country)
	if err != nil {
		return nil, errors.HandleError(log.Error, errors.New("ArtistRepository.FindArtistByID", "Error while reading data from db", err))
	}

	return artist, nil
}

// Save save an artist in the db
func (r ArtistRepository) Save(tx *sql.Tx, artist NewArtist) (*Artist, error) {
	log.Debug("ArtistRepository.Save - ", artist.String())

	result, err := tx.Exec("INSERT INTO artists (name, country) VALUES (?, ?)", artist.Name, artist.Country)
	if err != nil {
		return nil, errors.HandleError(log.Error, errors.New("ArtistRepository.Save", fmt.Sprintf("Error while saving artist %s", artist.String()), err))
	}

	id, _ := result.LastInsertId()
	return &Artist{ID: id, Name: artist.Name, Country: artist.Country}, nil
}

func (r ArtistRepository) FindArtistDiscography(id int) (*Discography, error) {
	log.Debugf("ArtistRepository.FindArtistDiscography - Artist ID = %d", id)

	rows, err := r.db.Query(
		"SELECT a.id, a.name, a.country, "+
			"r.id as r_id, r.title as r_title, r.year, r.genre, r.support, r.nb_support as r_nb_support, r.label, (select count(*) from tracks where id_record = r.id) as nb_tracks,"+
			"t.id as t_id, t.number, t.title as t_title, t.length, t.nb_support as t_nb_support "+
			"FROM artists a LEFT JOIN records r ON a.id = r.id_artist LEFT JOIN tracks t ON r.id = t.id_record "+
			"WHERE a.id = ? "+
			"ORDER BY r.year, r.id, t.number", id)
	if err != nil {
		return nil, errors.HandleError(log.Error, errors.New("ArtistRepository.FindArtistDiscography", "Database error", err))
	}
	defer rows.Close()

	return r.parseArtistDiscography(rows)
}

func (r ArtistRepository) parseArtistDiscography(rows *sql.Rows) (*Discography, error) {
	var discography = new(Discography)
	var records = make([]Record, 0)
	var record Record
	var tracks []*Track
	var count = 0

	for rows.Next() {
		var rId, rYear, rNbSupport, rNbTracks, tId, tNumber, tLength int
		var rTitle, rGenre, rSupport, rLabel, tTitle, tNbSupport string

		err := rows.Scan(&discography.ID, &discography.Name, &discography.Country,
			&rId, &rTitle, &rYear, &rGenre, &rSupport, &rNbSupport, &rLabel, &rNbTracks,
			&tId, &tNumber, &tTitle, &tLength, &tNbSupport)
		if err != nil {
			return nil, errors.HandleError(log.Error, errors.New("ArtistRepository.parseArtistDiscography", "Database error", err))
		}

		if rId == 1 { // new record
			record = Record{ID: rId, Title: rTitle, Year: &rYear, Genre: &rGenre, Support: &rSupport, NbSupport: &rNbSupport, Label: &rLabel}
			tracks = make([]*Track, 0)
		}

		tracks = append(tracks, &Track{ID: tId, Title: tTitle, Number: tNumber, Length: &tLength})

		if rNbTracks == tNumber { // last track of the current record
			record.Tracks = tracks
			records = append(records, record)
		}

		count++
	}

	if count == 0 {
		return nil, errors.HandleError(log.Error, errors.New("ArtistRepository.parseArtistDiscography", "Error while reading data from db", sql.ErrNoRows))
	}

	err := rows.Err()
	if err != nil {
		return nil, errors.HandleError(log.Error, errors.New("ArtistRepository.parseArtistDiscography", "Error while reading data from db", err))
	}

	discography.Records = records
	return discography, nil

}

// NewArtistRepository create a new ArtistRepository
func NewArtistRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{db}
}
