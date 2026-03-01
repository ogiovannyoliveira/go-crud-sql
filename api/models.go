package api

import (
	"encoding/json"
	"errors"

	uuid "github.com/satori/go.uuid"
)

type ID uuid.UUID

func (id ID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

func (id ID) MarshalJSON() ([]byte, error) {
	s := id.UUID().String()
	return json.Marshal(s)
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	u, err := uuid.FromString(s)
	if err != nil {
		return err
	}

	*id = ID(u)
	return nil
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

func (u User) Validate() error {
	if u.FirstName == "" {
		return errors.New("First name is required.")
	}
	if len(u.FirstName) < 2 || len(u.FirstName) > 20 {
		return errors.New("First name should have length between 2 and 20 characters.")
	}

	if u.LastName == "" {
		return errors.New("Last name is required.")
	}
	if len(u.LastName) < 2 || len(u.LastName) > 20 {
		return errors.New("Last name should have length between 2 and 20 characters.")
	}

	if u.Biography == "" {
		return errors.New("Biography is required.")
	}
	if len(u.Biography) < 20 || len(u.Biography) > 450 {
		return errors.New("Biography should have length between 20 and 450 characters.")
	}

	return nil
}

type Application struct {
	Data map[ID]User
}

type UserResponse struct {
	ID        ID     `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

func NewUserResponse(id ID, user User) UserResponse {
	return UserResponse{
		ID:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Biography: user.Biography,
	}
}

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}
