package repository

import (
	"context"
	"database/sql"

	. "library-api/internal/model"
)

type BookRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateBook(ctx context.Context, book *Book) error {
	query := "INSERT INTO books (title, author, year) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, book.Title, book.Author, book.Year)
	return err
}

func (r *BookRepository) GetBookByID(ctx context.Context, id uint) (*Book, error) {
	book := &Book{}
	query := "SELECT id, title, author, year FROM books WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]*Book, error) {
	query := "SELECT id, title, author, year FROM books"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		book := &Book{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *BookRepository) UpdateBook(ctx context.Context, book *Book) error {
	query := "UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4"
	_, err := r.db.ExecContext(ctx, query, book.Title, book.Author, book.Year, book.ID)
	return err
}

func (r *BookRepository) DeleteBook(ctx context.Context, id uint) error {
	query := "DELETE FROM books WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
