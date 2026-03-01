package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/ogiovannyoliveira/go-crud-sql/internal/api"
	"github.com/ogiovannyoliveira/go-crud-sql/internal/api/models"
	"github.com/ogiovannyoliveira/go-crud-sql/internal/store"
)

func main() {
	slog.Info("Starting server at port :8080.")

	if err := run(); err != nil {
		slog.Error("Something went wrong", "error", err.Error())
		return
	}

	slog.Info("Server is closed. All systems offline.")
}

func run() error {
	db, err := store.Initialize()
	if err != nil {
		slog.Error("Error while trying to connect to database.", "error", err)
		return err
	}
	defer db.Close()
	slog.Info("Database connected successfully.")

	repo := store.NewRepositories(db)
	slog.Info("Repositories initializes successfully.")

	app := models.Application{Data: map[models.ID]models.User{}}
	handler := api.NewHandler(app, repo)

	s := http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
