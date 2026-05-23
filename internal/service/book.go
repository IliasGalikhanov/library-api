package service

import (
	"context"
	"database/sql"

	. "library-api/internal/model"
	"library-api/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func New(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(ctx context.Context, book *Book) error {
	if book.Title == "" || book.Author == "" || book.Year <= 0 || book.Year > 2027 {
		return sql.ErrNoRows
	}

	return s.repo.CreateBook(ctx, book)
}

func (s *BookService) GetBook(ctx context.Context, id uint) (*Book, error) {
	if id == 0 {
		return nil, sql.ErrNoRows
	}
	return s.repo.GetBookByID(ctx, id)
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]*Book, error) {
	return s.repo.GetAllBooks(ctx)
}

func (s *BookService) UpdateBook(ctx context.Context, book *Book) error {
	if book.ID == 0 || book.Title == "" || book.Author == "" || book.Year <= 0 || book.Year > 2027 {
		return sql.ErrNoRows
	}

	return s.repo.UpdateBook(ctx, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id uint) error {
	if id == 0 {
		return sql.ErrNoRows
	}

	return s.repo.DeleteBook(ctx, id)
}
