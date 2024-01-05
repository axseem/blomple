package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func ConnectDB(dbPath string) (*Queries, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	return New(db), nil
}
