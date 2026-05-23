package service

import (
	"context"
	"database/sql"
	"fmt"

	"library-api/internal/auth"
	. "library-api/internal/model"
	"library-api/internal/repository"
)

type UserService struct {
	repo      *repository.UserRepository
	secretKey string
}

func NewUserService(repo *repository.UserRepository, secretKey string) *UserService {
	return &UserService{repo: repo, secretKey: secretKey}
}

func (s *UserService) Register(ctx context.Context, email, password, name string) (*User, error) {
	if email == "" || password == "" || name == "" {
		return nil, fmt.Errorf("email, password and name are required")
	}

	if len(password) < 6 {
		return nil, fmt.Errorf("password must be at least 6 characters")
	}

	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &User{
		Email:    email,
		Password: hashedPassword,
		Name:     name,
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return user, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %v", err)
	}

	if !auth.CheckPassword(password, user.Password) {
		return "", fmt.Errorf("invalid password")
	}

	token, err := auth.GenerateToken(user.ID, user.Email, s.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func (s *UserService) GetUser(ctx context.Context, id uint) (*User, error) {
	if id == 0 {
		return nil, sql.ErrNoRows
	}

	return s.repo.GetUserByID(ctx, id)
}
