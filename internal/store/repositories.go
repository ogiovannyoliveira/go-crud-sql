package store

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/ogiovannyoliveira/go-crud-in-memory/internal/api/models"
)

type repositories struct {
	db *sql.DB
}

type Repositories interface {
	SaveUser(ctx context.Context, user models.UserResponse) (models.UserResponse, error)
	UpdateUser(ctx context.Context, id models.ID, user models.User) (models.UserResponse, error)
	GetUsers(ctx context.Context, id models.ID) ([]models.UserResponse, error)
	GetUserByID(ctx context.Context, id models.ID) (models.UserResponse, error)
	DeleteUser(ctx context.Context, id models.ID) (models.UserResponse, error)
}

func NewRepositories(db *sql.DB) Repositories {
	slog.Info("Initializing repositories...")
	return repositories{db}
}

// SaveUser implements [Repositories].
func (r repositories) SaveUser(ctx context.Context, user models.UserResponse) (models.UserResponse, error) {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO users (id, first_name, last_name, biography) VALUES (?, ?, ?, ?)",
		user.ID, user.FirstName, user.LastName, user.Biography,
	)

	if err != nil {
		slog.Error("Could not insert user.", "error", err)
		return models.UserResponse{}, errors.New("Could not insert user.")
	}

	return r.GetUserByID(ctx, user.ID)
}

// DeleteUser implements [Repositories].
func (r repositories) DeleteUser(ctx context.Context, id models.ID) (models.UserResponse, error) {
	panic("unimplemented")
}

// GetUserByID implements [Repositories].
func (r repositories) GetUserByID(ctx context.Context, id models.ID) (models.UserResponse, error) {
	var user models.UserResponse
	row := r.db.QueryRowContext(
		ctx,
		"SELECT id, first_name, last_name, biography FROM users WHERE id = ?",
		id,
	)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Biography)
	if err != nil {
		slog.Error("Could not find user in the database.", "error", err)
		return models.UserResponse{}, errors.New("Could not find user.")
	}

	return user, nil
}

// GetUsers implements [Repositories].
func (r repositories) GetUsers(ctx context.Context, id models.ID) ([]models.UserResponse, error) {
	users := make([]models.UserResponse, 0)
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, first_name, last_name, biography FROM users",
	)

	if err != nil {
		slog.Error("Could not get users in the database.", "error", err)
		return []models.UserResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserResponse
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Biography)
		if err != nil {
			slog.Error("Could not scan user.", "error", err)
			return nil, errors.New("Could not scan user.")
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Something went wrong while trying to get users: ", "error", err)
		return nil, errors.New("Something went wrong while trying to get users.")
	}

	return users, nil
}

// UpdateUser implements [Repositories].
func (r repositories) UpdateUser(ctx context.Context, id models.ID, user models.User) (models.UserResponse, error) {
	panic("unimplemented")
}
