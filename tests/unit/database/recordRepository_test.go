package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/object-it/goserv/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Ce teste ne marche pas alors qu'il devrait. Est ce un bug dans slqmock ?
func testFindRecordById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	cols := []string{"id", "title", "year", "genre", "support", "nb_support", "label", "id_track", "number", "title", "length"}
	rows := sqlmock.NewRows(cols)
	rows.AddRow(15, "The Sciences", 2018, "Doom", "CD", 0, "Thrid Man", 156, 1, "The Sciences", 184)
	rows.AddRow(15, "The Sciences", 2018, "Doom", "CD", 0, "Thrid Man", 157, 2, "Marijuanaut's Theme", 400)
	rows.AddRow(15, "The Sciences", 2018, "Doom", "CD", 0, "Thrid Man", 158, 3, "Sonic Titan", 747)
	rows.AddRow(15, "The Sciences", 2018, "Doom", "CD", 0, "Thrid Man", 159, 4, "Antarcticans Thawed", 863)
	rows.AddRow(15, "The Sciences", 2018, "Doom", "CD", 0, "Thrid Man", 160, 5, "Giza Butler", 603)
	rows.AddRow(15, "The Sciences", 2018, "Doom", "CD", 0, "Thrid Man", 161, 6, "The Botanist", 387)
	mock.ExpectQuery(database.SelectRecordWithTracksByIdRecord).WithArgs(15).WillReturnRows(rows)

	repo := database.NewRecordRepository(db)
	rec, err := repo.FindRecordByID(15)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, rec)
	assert.Equal(t, 15, rec.ID)
	assert.Equal(t, "The Sciences", rec.Title)
	assert.Equal(t, 2018, rec.Year)
	assert.Equal(t, "Doom", rec.Genre)
	assert.Equal(t, "CD", rec.Support)
	assert.Equal(t, "Third Man", rec.Label)
	assert.Equal(t, 6, len(rec.Tracks))
	assert.Equal(t, "The Sciences", rec.Tracks[0].Title)
	assert.Equal(t, "Marijuanaut's Theme", rec.Tracks[1].Title)
	assert.Equal(t, "Sonic Titan", rec.Tracks[2].Title)
	assert.Equal(t, "Antarcticans Thawed", rec.Tracks[3].Title)
	assert.Equal(t, "Giza Butler", rec.Tracks[4].Title)
	assert.Equal(t, "The Botanist", rec.Tracks[5].Title)
}
