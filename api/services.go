package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	uuid "github.com/satori/go.uuid"
)

func Insert(app Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUser User

		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			SendJSON(w, Response{Error: "Could not parse body."}, http.StatusUnprocessableEntity)
			return
		}

		if err := newUser.Validate(); err != nil {
			SendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		id := ID(uuid.NewV4())
		app.Data[id] = newUser

		SendJSON(w, Response{
			Data:    NewUserResponse(id, newUser),
			Message: "User saved successfully.",
		}, http.StatusCreated)
	}
}

func FindAll(app Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := make([]UserResponse, 0, len(app.Data))

		for id, body := range app.Data {
			users = append(users, NewUserResponse(id, body))
		}

		SendJSON(w, Response{Data: users}, http.StatusOK)
	}
}

func FindByID(app Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		uid, err := uuid.FromString(idParam)
		if err != nil {
			SendJSON(w, Response{Error: "ID format is not a UUID."}, http.StatusBadRequest)
			return
		}

		user, ok := app.Data[ID(uid)]
		if !ok {
			SendJSON(w, Response{Error: "Could not find user."}, http.StatusNotFound)
			return
		}

		SendJSON(w, Response{Data: NewUserResponse(ID(uid), user)}, http.StatusOK)
	}
}

func Update(app Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		uid, err := uuid.FromString(idParam)
		if err != nil {
			SendJSON(w, Response{Error: "ID format is not a UUID."}, http.StatusBadRequest)
			return
		}

		var body User
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			SendJSON(w, Response{Error: "Could not parse body."}, http.StatusUnprocessableEntity)
			return
		}

		user, ok := app.Data[ID(uid)]
		if !ok {
			SendJSON(w, Response{Error: "Could not find user."}, http.StatusNotFound)
			return
		}

		user.FirstName = body.FirstName
		user.LastName = body.LastName
		user.Biography = body.Biography

		if err := user.Validate(); err != nil {
			SendJSON(w, Response{Error: err.Error()}, http.StatusBadRequest)
			return
		}

		app.Data[ID(uid)] = user
		SendJSON(w, Response{
			Data:    NewUserResponse(ID(uid), user),
			Message: "User updated successfully.",
		}, http.StatusOK)
	}
}

func Delete(app Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")

		uid, err := uuid.FromString(idParam)
		if err != nil {
			SendJSON(w, Response{Error: "ID format is not a UUID."}, http.StatusBadRequest)
			return
		}

		user, ok := app.Data[ID(uid)]
		if !ok {
			SendJSON(w, Response{Error: "Could not find user."}, http.StatusNotFound)
			return
		}

		delete(app.Data, ID(uid))
		SendJSON(w, Response{
			Data:    NewUserResponse(ID(uid), user),
			Message: "User deleted successfully.",
		}, http.StatusOK)
	}
}
