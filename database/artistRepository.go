package database

import (
	"database/sql"
	"fmt"
	"github.com/object-it/goserv/xerror"
	log "github.com/sirupsen/logrus"
)

type ArtistRepository struct {
	db *sql.DB
}

// NewArtistRepository créé un nouveau repository
func NewArtistRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{db}
}

// FindArtistByID recherche un artiste par son ID
func (r ArtistRepository) FindArtistByID(id int) (*Artist, error) {
	log.Debugf("ArtistRepository.FindArtistByID - ID = %d", id)

	row := r.db.QueryRow("SELECT id, name, country FROM artists WHERE id = ?", id)
	artist := new(Artist)
	err := row.Scan(&artist.ID, &artist.Name, &artist.Country)
	if err != nil {
		return nil, xerror.HandleError(log.Error, xerror.New("ArtistRepository.FindArtistByID", "Error while reading data from db", err))
	}

	return artist, nil
}

// Save sauvegarde (INSERT) un artiste en base de donnée
func (r ArtistRepository) Save(tx *sql.Tx, artist NewArtist) (int64, error) {
	log.Debugf("ArtistRepository.Save - %v", artist)

	result, err := tx.Exec("INSERT INTO artists (name, country) VALUES (?, ?)", artist.Name, artist.Country)
	if err != nil {
		return -1, xerror.HandleError(log.Error, xerror.New("ArtistRepository.Save", fmt.Sprintf("Error while saving artist %v", artist), err))
	}

	return result.LastInsertId() // err is always nil
}

func (r ArtistRepository) Delete(tx *sql.Tx, id int) error {
	log.Debugf("ArtistRepository.Delete - ID = %d", id)

	if _, err := tx.Exec("DELETE FROM artists WHERE id = ?", id); err != nil {
		return xerror.HandleError(log.Error, xerror.New("ArtistRepository.Delete", "Database error", err))
	}

	return nil
}

// FindArtistDiscography charge la discographie d'un artiste
func (r ArtistRepository) FindArtistDiscography(id int) (*Discography, error) {
	log.Debugf("ArtistRepository.FindArtistDiscography - Artist ID = %d", id)

	rows, err := r.db.Query("SELECT a.id, a.name, a.country, "+
		"r.id as r_id, r.title as r_title, r.year, r.genre, r.support, r.nb_support as r_nb_support, r.label, (select count(*) from tracks where id_record = r.id) as nb_tracks,"+
		"t.id as t_id, t.number, t.title as t_title, t.length, t.nb_support as t_nb_support "+
		"FROM artists a INNER JOIN records r ON a.id = r.id_artist INNER JOIN tracks t ON r.id = t.id_record "+
		"WHERE a.id = ? "+
		"ORDER BY r.year, r.id, t.number", id)
	if err != nil {
		return nil, xerror.HandleError(log.Error, xerror.New("ArtistRepository.FindArtistDiscography", "Database error", err))
	}
	defer rows.Close()

	return r.parseArtistDiscography(rows)
}

func (r ArtistRepository) parseArtistDiscography(rows *sql.Rows) (*Discography, error) {
	var discography = Discography{Records: make([]Record, 0)}
	var record *Record

	for rows.Next() {
		var rId, rNbTracks, tId, tNumber int64
		var rTitle, tTitle string
		var rGenre, rSupport, rLabel NullString
		var rYear, rNbSupport, tLength, tNbSupport NullInt64

		err := rows.Scan(&discography.ID, &discography.Name, &discography.Country,
			&rId, &rTitle, &rYear, &rGenre, &rSupport, &rNbSupport, &rLabel, &rNbTracks,
			&tId, &tNumber, &tTitle, &tLength, &tNbSupport)
		if err != nil {
			return nil, xerror.HandleError(log.Error, xerror.New("ArtistRepository.parseArtistDiscography", "Database error", err))
		}

		if record == nil || record.ID != rId {
			record = &Record{ID: rId, Title: rTitle, Year: rYear, Genre: rGenre,
				Support: rSupport, NbSupport: rNbSupport, Label: rLabel, Tracks: make([]Track, 0)}
		}

		record.Tracks = append(record.Tracks, Track{ID: tId, Title: tTitle, Number: tNumber, Length: tLength})

		if rNbTracks == tNumber {
			discography.Records = append(discography.Records, *record)
		}
	}

	err := rows.Err()
	if err != nil {
		return nil, xerror.HandleError(log.Error, xerror.New("ArtistRepository.parseArtistDiscography", "Error while reading data from db", err))
	}

	//noinspection ALL
	if record == nil {
		return nil, xerror.HandleError(log.Error, xerror.New("ArtistRepository.parseArtistDiscography", "Error while reading data from db", sql.ErrNoRows))
	}

	return &discography, nil
}
