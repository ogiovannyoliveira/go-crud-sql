package store

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
)

func Initialize() (*sql.DB, error) {
	slog.Info("Connecting to the database...")

	conn := "user=giovanny dbname=go-crud-sql password=secret sslmode=disabled"
	db, err := sql.Open("postgres", conn)

	if err != nil {
		return nil, fmt.Errorf("Could not connect to database: %w", err)
	}

	slog.Info("Connection to database has been established successfully.")
	return db, nil
}
