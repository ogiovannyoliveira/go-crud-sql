package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandler(app Application) http.Handler {
	r := chi.NewMux()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		SendJSON(w, Response{Data: "System is healthy and running"}, http.StatusOK)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", Insert(app))
			r.Get("/", FindAll(app))
			r.Get("/{id}", FindByID(app))
			r.Put("/{id}", Update(app))
			r.Delete("/{id}", Delete(app))
		})
	})

	return r
}
