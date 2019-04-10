package database

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/object-it/goserv/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetArtistByID_ShouldReturnAnArtist_WhenArtistIsInDatabase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	cols := []string{"id", "name", "country"}
	rows := sqlmock.NewRows(cols)
	rows = rows.AddRow(1, "Sleep", "USA")
	mock.ExpectQuery(database.SelectArtistById).WithArgs(1).WillReturnRows(rows)

	repo := database.NewArtistRepository(db)
	artist, err := repo.FindArtistByID(1)

	assert.Nil(t, err, err)
	assert.NotNil(t, artist, "Artist should not be null !!!")
	assert.Equal(t, int64(1), artist.ID, "Artist's ID should be equal to 1")
	assert.Equal(t, "Sleep", artist.Name, "Artist's name should be Sleep")
	assert.Equal(t, "USA", artist.Country.String, "Artist's country should be USA")
}

func TestGetArtistByID_ShouldThrowAnError_WhenArtistIsNotInDatabase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(database.SelectArtistById).WithArgs(sqlmock.AnyArg()).WillReturnError(sql.ErrNoRows)

	repo := database.NewArtistRepository(db)
	artist, err := repo.FindArtistByID(2)

	assert.NotNil(t, err, "There should have been an error")
	assert.Contains(t, err.Error(), sql.ErrNoRows.Error())
	assert.Nil(t, artist, "Artist should be null !!!")
}
