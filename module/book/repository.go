package book

import (
	"book-backend/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

type mrRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) domain.BookRepository {
	return &mrRepository{db}
}

func (r *mrRepository) FindAll(ctx context.Context) ([]domain.Book, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT * FROM Books")
	if err != nil {
		return nil, err
	}

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		var createdAt, updatedAt string
		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Price,
			&book.ImageUrl,
			&book.Author,
			&book.Genre,
			&book.YearOfPublication,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		book.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}
		book.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *mrRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM Books WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r *mrRepository) FindByID(ctx context.Context, id int64) (*domain.Book, error) {
	var book domain.Book
	var createdAt, updatedAt string
	err := r.DB.QueryRowContext(ctx, "SELECT * FROM Books WHERE id = ?", id).
		Scan(
			&book.ID,
			&book.Title,
			&book.Price,
			&book.ImageUrl,
			&book.Author,
			&book.Genre,
			&book.YearOfPublication,
			&createdAt,
			&updatedAt,
		)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	book.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		return nil, err
	}
	book.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *mrRepository) Insert(ctx context.Context, book *domain.Book) error {
	_, err := r.DB.ExecContext(
		ctx,
		"INSERT INTO Books (title, price, ImageUrl, author, genre, YearOfPublication, CreatedAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		book.Title,
		book.Price,
		book.ImageUrl,
		book.Author,
		book.Genre,
		book.YearOfPublication,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *mrRepository) Update(ctx context.Context, id int64, book *domain.Book) error {
	_, err := r.DB.ExecContext(
		ctx,
		"UPDATE Books SET title = ?, price = ?, ImageUrl = ?, author = ?, genre = ?, YearOfPublication = ?, UpdatedAt = ? WHERE id = ?",
		book.Title,
		book.Price,
		book.ImageUrl,
		book.Author,
		book.Genre,
		book.YearOfPublication,
		time.Now(),
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *mrRepository) FindByTitleAndAuthor(ctx context.Context, title, author string) (*domain.Book, error) {
	var book domain.Book
	var createdAt, updatedAt string
	err := r.DB.QueryRowContext(ctx, "SELECT * FROM Books WHERE title = ? AND author = ?", title, author).
		Scan(
			&book.ID,
			&book.Title,
			&book.Price,
			&book.ImageUrl,
			&book.Author,
			&book.Genre,
			&book.YearOfPublication,
			&createdAt,
			&updatedAt,
		)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	book.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		return nil, err
	}
	book.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAt)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
