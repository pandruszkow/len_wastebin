package storage

import (
	"database/sql"
	"time"
)

type Paste struct {
	ID         string `json:"id"` // Ignored when creating
	Title      string `json:"title"`
	Body       string `json:"body"`
	CreateTime int64  `json:"createTime"` // Ignored when creating
	DeleteTime int64  `json:"deleteTime"`
	OneUse     bool   `json:"oneUse"`
	Syntax     string `json:"syntax"`

	Author      string `json:"author"`
	AuthorEmail string `json:"authorEmail"`
	AuthorURL   string `json:"authorURL"`
}

func (dbInfo DB) PasteAdd(paste Paste) (string, error) {
	// Open DB
	db, err := dbInfo.openDB()
	if err != nil {
		return paste.ID, err
	}
	defer db.Close()

	// Generate ID
	paste.ID, err = genTokenCrypto(8)
	if err != nil {
		return paste.ID, err
	}

	// Set paste create time
	paste.CreateTime = time.Now().Unix()

	// Check delete time
	if paste.DeleteTime < 0 {
		paste.DeleteTime = 0
	}

	// Add
	_, err = db.Exec(
		`INSERT INTO pastes (id, title, body, syntax, create_time, delete_time, one_use, author, author_email, author_url) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		paste.ID, paste.Title, paste.Body, paste.Syntax, paste.CreateTime, paste.DeleteTime, paste.OneUse, paste.Author, paste.AuthorEmail, paste.AuthorURL,
	)
	if err != nil {
		return paste.ID, err
	}

	return paste.ID, nil
}

func (dbInfo DB) PasteDelete(id string) error {
	// Open DB
	db, err := dbInfo.openDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Delete
	result, err := db.Exec(
		`DELETE FROM pastes WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	// Check result
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFoundID
	}

	return nil
}

func (dbInfo DB) PasteGet(id string) (Paste, error) {
	var paste Paste

	// Open DB
	db, err := dbInfo.openDB()
	if err != nil {
		return paste, err
	}
	defer db.Close()

	// Make query
	row := db.QueryRow(
		`SELECT id, title, body, syntax, create_time, delete_time, one_use, author, author_email, author_url FROM pastes WHERE id = $1`,
		id,
	)

	// Read query
	err = row.Scan(&paste.ID, &paste.Title, &paste.Body, &paste.Syntax, &paste.CreateTime, &paste.DeleteTime, &paste.OneUse, &paste.Author, &paste.AuthorEmail, &paste.AuthorURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return paste, ErrNotFoundID
		}

		return paste, err
	}

	// Check paste expiration
	if paste.DeleteTime < time.Now().Unix() && paste.DeleteTime > 0 {
		// Delete expired paste
		_, err = db.Exec(
			`DELETE FROM pastes WHERE id = $1`,
			paste.ID,
		)
		if err != nil {
			return Paste{}, err
		}

		// Return ErrNotFound
		return Paste{}, ErrNotFoundID
	}

	return paste, nil
}

func (dbInfo DB) PasteDeleteExpired() (int64, error) {
	// Open DB
	db, err := dbInfo.openDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	// Delete
	result, err := db.Exec(
		`DELETE FROM pastes WHERE (delete_time < $1) AND (delete_time > 0)`,
		time.Now().Unix(),
	)
	if err != nil {
		return 0, err
	}

	// Check result
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, nil
}
