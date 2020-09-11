package database

import (
	"database/sql"
	"log"

	// Load sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// CreateConn Create connection to local SQLite3 DB
func CreateConn(dbpath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
