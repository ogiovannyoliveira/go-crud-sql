package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/ogiovannyoliveira/go-crud-in-memory/api"
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
	app := api.Application{Data: map[api.ID]api.User{}}
	handler := api.NewHandler(app)

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
