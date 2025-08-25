package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB(filepath string) error {
	var err error
	db, err = sql.Open("sqlite", filepath)
	if err != nil {
		return err
	}

	// Create table if not exists
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		original_name TEXT,
		stored_name TEXT,
		size INTEGER,
		uploaded_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(sqlStmt)
	return err
}
