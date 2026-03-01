package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/ogiovannyoliveira/go-crud-in-memory/internal/api"
	"github.com/ogiovannyoliveira/go-crud-in-memory/internal/api/models"
	"github.com/ogiovannyoliveira/go-crud-in-memory/internal/store"
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

	repo := store.NewRepositories(db)

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
