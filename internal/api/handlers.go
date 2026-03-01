package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogiovannyoliveira/go-crud-sql/internal/api/models"
)

func NewHandler(services Services) http.Handler {
	r := chi.NewMux()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		SendJSON(w, models.Response{Data: "System is healthy and running"}, http.StatusOK)
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", services.Insert())
			r.Get("/", services.FindAll())
			r.Get("/{id}", services.FindByID())
			r.Put("/{id}", services.Update())
			r.Delete("/{id}", services.Delete())
		})
	})

	return r
}
