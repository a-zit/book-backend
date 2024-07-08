package domain

import (
	"context"
	"time"
)

type Book struct {
	ID                int64     `json:"id"                  gorm:"primary_key,column:ID"`
	Title             string    `json:"title"               gorm:"column:Title"`
	Price             float64   `json:"price"               gorm:"column:Price"`
	ImageUrl          string    `json:"image_url"           gorm:"column:ImageUrl"`
	Author            string    `json:"author"              gorm:"column:Author"`
	Genre             string    `json:"genre"               gorm:"column:Genre"`
	YearOfPublication int       `json:"year_of_publication" gorm:"column:YearOfPublication"`
	CreatedAt         time.Time `json:"created_at"          gorm:"column:CreatedAt"`
	UpdatedAt         time.Time `json:"updated_at"          gorm:"column:UpdatedAt"`
}

type CreateBookRequest struct {
	Title             string  `json:"title"               validate:"required"`
	Price             float64 `json:"price"               validate:"required"`
	ImageUrl          string  `json:"image_url"           validate:""`
	Author            string  `json:"author"              validate:"required"`
	Genre             string  `json:"genre"               validate:"required"`
	YearOfPublication int     `json:"year_of_publication" validate:"required,year"`
}

type EditBookRequest struct {
	Title             string  `json:"title"               validate:"required"`
	Price             float64 `json:"price"               validate:"required"`
	ImageUrl          string  `json:"image_url"           validate:""`
	Author            string  `json:"author"              validate:"required"`
	Genre             string  `json:"genre"               validate:"required"`
	YearOfPublication int     `json:"year_of_publication" validate:"required,year"`
}

type BookRepository interface {
	FindAll(ctx context.Context) ([]Book, error)
	FindByID(ctx context.Context, id int64) (*Book, error)
	FindByTitleAndAuthor(ctx context.Context, title, author string) (*Book, error)
	Insert(ctx context.Context, book *Book) error
	Update(ctx context.Context, id int64, book *Book) error
	Delete(ctx context.Context, id int64) error
}

type BookUseCase interface {
	GetAll(ctx context.Context) ([]Book, error)
	GetByID(ctx context.Context, id int64) (*Book, error)
	Create(ctx context.Context, request CreateBookRequest) error
	Edit(ctx context.Context, request EditBookRequest, id int64) error
	Delete(ctx context.Context, id int64) error
}
