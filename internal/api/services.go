package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogiovannyoliveira/go-crud-sql/internal/api/models"
	"github.com/ogiovannyoliveira/go-crud-sql/internal/store"
	uuid "github.com/satori/go.uuid"
)

type services struct {
	repo store.Repositories
}

type Services interface {
	Insert() http.HandlerFunc
	FindAll() http.HandlerFunc
	FindByID() http.HandlerFunc
	Update() http.HandlerFunc
	Delete() http.HandlerFunc
}

func NewServices(repo store.Repositories) Services {
	slog.Info("Initializing services...")
	return services{repo}
}

// Insert implements [Services].
func (s services) Insert() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser models.User

		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			slog.Error("Could not parse body:", "error", err)
			SendJSON(w, models.Response{Error: "Could not parse body."}, http.StatusUnprocessableEntity)
			return
		}

		if err := newUser.Validate(); err != nil {
			slog.Error("Could validate user:", "error", err)
			SendJSON(w, models.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		id := models.ID(uuid.NewV4())
		user, err := s.repo.SaveUser(r.Context(), models.NewUserResponse(id, newUser))
		if err != nil {
			SendJSON(w, models.Response{Error: err.Error()}, http.StatusInternalServerError)
			return
		}

		SendJSON(w, models.Response{
			Data:    user,
			Message: "User saved successfully.",
		}, http.StatusCreated)
	}
}

// FindAll implements [Services].
func (s services) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.repo.GetUsers(r.Context())
		if err != nil {
			SendJSON(w, models.Response{Error: err.Error()}, http.StatusInternalServerError)
			return
		}

		SendJSON(w, models.Response{Data: users}, http.StatusOK)
	}
}

// FindByID implements [Services].
func (s services) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		uid, err := uuid.FromString(idParam)
		if err != nil {
			slog.Error("Could not parse ID:", "error", err)
			SendJSON(w, models.Response{Error: "ID format is not a UUID."}, http.StatusBadRequest)
			return
		}

		user, err := s.repo.GetUserByID(r.Context(), models.ID(uid))
		if err != nil {
			slog.Error("Could not get user:", "error", err)
			SendJSON(w, models.Response{Error: "Could not find user."}, http.StatusNotFound)
			return
		}

		SendJSON(w, models.Response{Data: user}, http.StatusOK)
	}
}

// Update implements [Services].
func (s services) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		uid, err := uuid.FromString(idParam)
		if err != nil {
			SendJSON(w, models.Response{Error: "ID format is not a UUID."}, http.StatusBadRequest)
			return
		}

		var body models.User
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			SendJSON(w, models.Response{Error: "Could not parse body."}, http.StatusUnprocessableEntity)
			return
		}

		if err := body.Validate(); err != nil {
			SendJSON(w, models.Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		user, err := s.repo.GetUserByID(r.Context(), models.ID(uid))
		if err != nil {
			SendJSON(w, models.Response{Error: "Could not find user."}, http.StatusNotFound)
			return
		}

		updated, err := s.repo.UpdateUser(r.Context(), models.ID(uid), body)
		if !updated || err != nil {
			SendJSON(w, models.Response{Error: "Could not update user."}, http.StatusInternalServerError)
			return
		}

		user.FirstName = body.FirstName
		user.LastName = body.LastName
		user.Biography = body.Biography

		SendJSON(w, models.Response{
			Data:    user,
			Message: "User updated successfully.",
		}, http.StatusOK)
	}
}

// Delete implements [Services].
func (s services) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		uid, err := uuid.FromString(idParam)
		if err != nil {
			SendJSON(w, models.Response{Error: "ID format is not a UUID."}, http.StatusBadRequest)
			return
		}

		user, err := s.repo.GetUserByID(r.Context(), models.ID(uid))
		if err != nil {
			SendJSON(w, models.Response{Error: "Could not find user."}, http.StatusNotFound)
			return
		}

		deleted, err := s.repo.DeleteUser(r.Context(), models.ID(uid))
		if !deleted || err != nil {
			SendJSON(w, models.Response{Error: "Could not delete user."}, http.StatusInternalServerError)
			return
		}

		SendJSON(w, models.Response{
			Data:    user,
			Message: "User deleted successfully.",
		}, http.StatusOK)
	}
}
