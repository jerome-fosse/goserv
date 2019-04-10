package database

const (
	SelectArtistById            = "SELECT id, name, country FROM artists WHERE id = ?"
	InsertIntoArtists           = "INSERT INTO artists (name, country) VALUES (?, ?)"
	DeleteArtistById            = "DELETE FROM artists WHERE id = ?"
	SelectArtistWithDiscography = `SELECT a.id, a.name, a.country,
 r.id as r_id, r.title as r_title, r.year, r.genre, r.support, r.nb_support as r_nb_support, r.label, (select count(*) from tracks where id_record = r.id) as nb_tracks,
 t.id as t_id, t.number, t.title as t_title, t.length, t.nb_support as t_nb_support
 FROM artists a INNER JOIN records r ON a.id = r.id_artist INNER JOIN tracks t ON r.id = t.id_record
 WHERE a.id = ?
 ORDER BY r.year, r.id, t.number`

	InsertIntoRecords                = "INSERT INTO records(title, id_artist, year, genre, support, nb_support, label) VALUES(?, ?, ?, ?, ?, ?, ?)"
	InsertIntoTracks                 = "INSERT INTO tracks (id_record, number, title, length) VALUES(?, ?, ?, ?)"
	DeleteRecordById                 = "DELETE FROM records WHERE id = ?"
	SelectRecordWithTracksByIdRecord = "SELECT rec.id, rec.title, rec.year, rec.genre, rec.support, rec.nb_support, rec.label, " +
		"tra.id as id_track, tra.number, tra.title, tra.length " +
		"FROM records rec LEFT JOIN tracks tra on rec.id = tra.id_record " +
		"WHERE rec.id = ? " +
		"ORDER BY tra.number ASC"
)
