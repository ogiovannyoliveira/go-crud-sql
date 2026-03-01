package store

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
)

func Initialize() (*sql.DB, error) {
	slog.Info("Connecting to the database...")

	conn := "user=giovanny dbname=go-crud-sql password=secret sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		slog.Error("Could not connect to database.")
		return nil, fmt.Errorf("Could not connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		slog.Error("Could not validate connection to the database.")
		return nil, fmt.Errorf("Could not validate connection to the database: %w", err)
	}
	slog.Info("Connection to the database is healthy.")

	return db, nil
}
