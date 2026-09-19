package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func InitDB(dataSourceName string) (*sql.DB, error) {
	err := os.Mkdir("data", 0755)

	if err != nil {
		return nil, fmt.Errorf("Failed to create data directory: %w", err)
	}

	db, err := sql.Open("sqlite", dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("Failed to open the database: %w", err)
	}

	_, err = db.Exec("PRAGMA foreign_keys = ON;")

	if err != nil {
		return nil, fmt.Errorf("Failed to enable foreign keys: %w", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	return db, nil
}
