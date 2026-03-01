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

func Insert(repo store.Repositories) http.HandlerFunc {
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
		user, err := repo.SaveUser(r.Context(), models.NewUserResponse(id, newUser))
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

func FindAll(repo store.Repositories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repo.GetUsers(r.Context())
		if err != nil {
			SendJSON(w, models.Response{Error: err.Error()}, http.StatusInternalServerError)
			return
		}

		SendJSON(w, models.Response{Data: users}, http.StatusOK)
	}
}

func FindByID(repo store.Repositories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		uid, err := uuid.FromString(idParam)
		if err != nil {
			slog.Error("Could not parse ID:", "error", err)
			SendJSON(w, models.Response{Error: "ID format is not a UUID."}, http.StatusBadRequest)
			return
		}

		user, err := repo.GetUserByID(r.Context(), models.ID(uid))
		if err != nil {
			slog.Error("Could not get user:", "error", err)
			SendJSON(w, models.Response{Error: "Could not find user."}, http.StatusNotFound)
			return
		}

		SendJSON(w, models.Response{Data: user}, http.StatusOK)
	}
}

func Update(repo store.Repositories) http.HandlerFunc {
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

		user, err := repo.GetUserByID(r.Context(), models.ID(uid))
		if err != nil {
			SendJSON(w, models.Response{Error: "Could not find user."}, http.StatusNotFound)
			return
		}

		updated, err := repo.UpdateUser(r.Context(), models.ID(uid), body)
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

func Delete(repo store.Repositories) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		uid, err := uuid.FromString(idParam)
		if err != nil {
			SendJSON(w, models.Response{Error: "ID format is not a UUID."}, http.StatusBadRequest)
			return
		}

		user, err := repo.GetUserByID(r.Context(), models.ID(uid))
		if err != nil {
			SendJSON(w, models.Response{Error: "Could not find user."}, http.StatusNotFound)
			return
		}

		deleted, err := repo.DeleteUser(r.Context(), models.ID(uid))
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
