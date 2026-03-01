package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogiovannyoliveira/go-crud-sql/internal/api/models"
	"github.com/ogiovannyoliveira/go-crud-sql/internal/store"
)

func NewHandler(app models.Application, repo store.Repositories) http.Handler {
	r := chi.NewMux()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		SendJSON(w, models.Response{Data: "System is healthy and running"}, http.StatusOK)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", Insert(repo))
			r.Get("/", FindAll(repo))
			r.Get("/{id}", FindByID(repo))
			r.Put("/{id}", Update(repo))
			r.Delete("/{id}", Delete(repo))
		})
	})

	return r
}
