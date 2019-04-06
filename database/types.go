package database

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}
type NullInt64 struct {
	sql.NullInt64
}
type NullFloat64 struct {
	sql.NullFloat64
}
type NullBool struct {
	sql.NullBool
}

func (s NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}

	return []byte("null"), nil
}

func (i NullInt64) MarshalJSON() ([]byte, error) {
	if i.Valid {
		return json.Marshal(i.Int64)
	}

	return []byte("null"), nil
}

func (f NullFloat64) MarshalJSON() ([]byte, error) {
	if f.Valid {
		return json.Marshal(f.Float64)
	}

	return []byte("null"), nil
}

func (b NullBool) MarshalJSON() ([]byte, error) {
	if b.Valid {
		return json.Marshal(b.Bool)
	}

	return []byte("null"), nil
}

func newNullString(s string) NullString {
	return NullString{sql.NullString{String: s, Valid: true}}
}
