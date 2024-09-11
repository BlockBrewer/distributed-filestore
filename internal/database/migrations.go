package database

import "database/sql"

func MigrateTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS file_chunks (
			id SERIAL PRIMARY KEY,
			file_id TEXT,
			data BYTEA,
			part INT
		);
		CREATE TABLE IF NOT EXISTS file_metadata (
			id TEXT PRIMARY KEY,
			name TEXT,
			size BIGINT
		);
	`)
	return err
}
