package book

import (
	"context"

	domain "book-backend/domain"
	"book-backend/pkg/validator"
)

type useCase struct {
	bookRepo domain.BookRepository
}

func NewBookUseCase(bookRepo domain.BookRepository) domain.BookUseCase {
	return &useCase{
		bookRepo: bookRepo,
	}
}

func (u *useCase) GetAll(ctx context.Context) ([]domain.Book, error) {
	books, err := u.bookRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *useCase) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	book, err := u.bookRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, domain.ErrBookNotFound
	}

	return book, nil
}

func (u *useCase) Delete(ctx context.Context, id int64) error {
	book, err := u.bookRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if book == nil {
		return domain.ErrBookNotFound
	}

	err = u.bookRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *useCase) Create(ctx context.Context, request domain.CreateBookRequest) error {
	// struct validator
	errs := validator.StructValidator(request)
	if errs != nil {
		return domain.ErrValidationFailed.SetMessage(errs)
	}
	// check exist book by author and title
	book, err := u.bookRepo.FindByTitleAndAuthor(ctx, request.Title, request.Author)
	if err != nil {
		return err
	}
	if book != nil {
		return domain.ErrBookAlreadyExist
	}
	// create book
	err = u.bookRepo.Insert(ctx, &domain.Book{
		Title:             request.Title,
		Price:             request.Price,
		ImageUrl:          request.ImageUrl,
		Author:            request.Author,
		Genre:             request.Genre,
		YearOfPublication: request.YearOfPublication,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *useCase) Edit(ctx context.Context, request domain.EditBookRequest, id int64) error {
	// struct validator
	errs := validator.StructValidator(request)
	if errs != nil {
		return domain.ErrValidationFailed.SetMessage(errs)
	}
	// find book
	book, err := u.bookRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if book == nil {
		return domain.ErrBookNotFound
	}
	// check exist book by author and title
	if request.Title != book.Title || request.Author != book.Author {
		book, err = u.bookRepo.FindByTitleAndAuthor(ctx, request.Title, request.Author)
		if err != nil {
			return err
		}
		if book != nil {
			return domain.ErrBookAlreadyExist
		}
	}
	// update book
	err = u.bookRepo.Update(ctx, id, &domain.Book{
		Title:             request.Title,
		Price:             request.Price,
		ImageUrl:          request.ImageUrl,
		Author:            request.Author,
		Genre:             request.Genre,
		YearOfPublication: request.YearOfPublication,
	})
	if err != nil {
		return err
	}

	return nil
}
