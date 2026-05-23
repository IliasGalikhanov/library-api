package repository

import (
	"context"
	"database/sql"
	"errors"

	. "library-api/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *User) error {
	query := "INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Password, user.Name).Scan(&user.ID)
	return err
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	query := "SELECT id, email, password, name FROM users WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uint) (*User, error) {
	user := &User{}
	query := "SELECT id, email, password, name FROM users WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.Password, &user.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // User not found
		}
		return nil, err
	}
	return user, nil
}
