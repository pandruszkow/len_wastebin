package storage

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrNotFoundID = errors.New("db: could not find ID")
)

type DB struct {
	DriverName     string
	DataSourceName string
}

func (dbInfo DB) openDB() (*sql.DB, error) {
	db, err := sql.Open(dbInfo.DriverName, dbInfo.DataSourceName)
	return db, err
}

func (dbInfo DB) InitDB() error {
	// Open DB
	db, err := dbInfo.openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Create tables
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS pastes (
			id          TEXT    PRIMARY KEY,
			title       TEXT    NOT NULL,
			body        TEXT    NOT NULL,
			syntax      TEXT    NOT NULL,
			create_time INTEGER NOT NULL,
			delete_time INTEGER NOT NULL,
			one_use     BOOL    NOT NULL
		);
	`)
	if err != nil {
		return err
	}

	// Crutch for SQLite3
	if dbInfo.DriverName == "sqlite3" {
		_, err = db.Exec(`ALTER TABLE pastes ADD COLUMN author       TEXT NOT NULL DEFAULT ''`)
		if err != nil {
			if err.Error() != "duplicate column name: author" {
				return err
			}
		}

		_, err = db.Exec(`ALTER TABLE pastes ADD COLUMN author_email TEXT NOT NULL DEFAULT ''`)
		if err != nil {
			if err.Error() != "duplicate column name: author_email" {
				return err
			}
		}

		_, err = db.Exec(`ALTER TABLE pastes ADD COLUMN author_url TEXT NOT NULL DEFAULT ''`)
		if err != nil {
			if err.Error() != "duplicate column name: author_url" {
				return err
			}
		}

		// Normal SQL for all other DBs
	} else {
		_, err = db.Exec(`
			ALTER TABLE pastes ADD COLUMN IF NOT EXISTS author       TEXT NOT NULL DEFAULT '';
			ALTER TABLE pastes ADD COLUMN IF NOT EXISTS author_email TEXT NOT NULL DEFAULT '';
			ALTER TABLE pastes ADD COLUMN IF NOT EXISTS author_url   TEXT NOT NULL DEFAULT '';
		`)
		if err != nil {
			return err
		}
	}

	return nil
}
