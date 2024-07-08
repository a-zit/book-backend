package book_suite

import (
	"book-backend/domain"
	"book-backend/domain/mocks"
	"book-backend/module/book"
	"book-backend/pkg/validator"
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BookSuite struct {
	suite.Suite
	useCase  domain.BookUseCase
	bookRepo *mocks.BookRepository
}

func (s *BookSuite) SetupTest() {
	validator.InitValidator()
	s.bookRepo = mocks.NewBookRepository(s.T())
	s.useCase = book.NewBookUseCase(s.bookRepo)
}

func (s *BookSuite) Teardown() {
	// clean suite
}

func (s *BookSuite) TestGetAll() {
	s.bookRepo.On("FindAll", context.Background()).Return([]domain.Book{}, nil)
	res, err := s.useCase.GetAll(context.Background())
	s.NoError(err)
	s.NotNil(res)
}

func (s *BookSuite) TestGetByIDButNotFound() {
	s.bookRepo.On("FindByID", context.Background(), int64(1)).Return(nil, nil)
	res, err := s.useCase.GetByID(context.Background(), int64(1))
	s.Error(err)
	s.Nil(res)
}

func (s *BookSuite) TestGetByID() {
	s.bookRepo.On("FindByID", context.Background(), int64(1)).Return(&domain.Book{}, nil)
	res, err := s.useCase.GetByID(context.Background(), int64(1))
	s.NoError(err)
	s.NotNil(res)
}

func (s *BookSuite) TestDeleteButNotFound() {
	s.bookRepo.On("FindByID", context.Background(), int64(1)).Return(nil, nil)
	err := s.useCase.Delete(context.Background(), int64(1))
	s.Error(err)
}

func (s *BookSuite) TestDelete() {
	s.bookRepo.On("FindByID", context.Background(), int64(1)).Return(&domain.Book{}, nil)
	s.bookRepo.On("Delete", context.Background(), int64(1)).Return(nil)
	err := s.useCase.Delete(context.Background(), int64(1))
	s.NoError(err)
}

func (s *BookSuite) TestCreateButErrorValidation() {
	req := domain.CreateBookRequest{
		Title: "Title",
		Price: 1000,
	}
	err := s.useCase.Create(context.Background(), req)
	s.Error(err)
	s.Contains(err.Error(), domain.ErrValidationFailed.Code)
}

func (s *BookSuite) TestCreateButErrorAlreadyExist() {
	req := domain.CreateBookRequest{
		Title:             "Title",
		Price:             1000,
		ImageUrl:          "https://image.com",
		Author:            "Author",
		Genre:             "Genre",
		YearOfPublication: 2021,
	}

	s.bookRepo.On("FindByTitleAndAuthor", context.Background(), req.Title, req.Author).
		Return(&domain.Book{}, nil)
	err := s.useCase.Create(context.Background(), req)
	s.Error(err)
	s.Equal(domain.ErrBookAlreadyExist, err)
}

func (s *BookSuite) TestCreate() {
	req := domain.CreateBookRequest{
		Title:             "Title",
		Price:             1000,
		ImageUrl:          "https://image.com",
		Author:            "Author",
		Genre:             "Genre",
		YearOfPublication: 2021,
	}

	s.bookRepo.On("FindByTitleAndAuthor", context.Background(), req.Title, req.Author).
		Return(nil, nil)
	s.bookRepo.On("Insert", context.Background(), mock.AnythingOfType("*domain.Book")).Return(nil)
	err := s.useCase.Create(context.Background(), req)
	s.NoError(err)
}

func (s *BookSuite) TestEditButErrorValidationFailed() {
	req := domain.EditBookRequest{
		Title: "Title",
		Price: 1000,
	}
	err := s.useCase.Edit(context.Background(), req, int64(1))
	s.Error(err)
	s.Contains(err.Error(), domain.ErrValidationFailed.Code)
}

// Not found
func (s *BookSuite) TestEditButErrorNotFound() {
	req := domain.EditBookRequest{
		Title:             "Title",
		Price:             1000,
		ImageUrl:          "https://image.com",
		Author:            "Author",
		Genre:             "Genre",
		YearOfPublication: 2021,
	}
	s.bookRepo.On("FindByID", context.Background(), int64(1)).Return(nil, nil)
	err := s.useCase.Edit(context.Background(), req, int64(1))
	s.Error(err)
	s.Equal(domain.ErrBookNotFound, err)
}

// Already exist
func (s *BookSuite) TestEditButErrorAlreadyExist() {
	req := domain.EditBookRequest{
		Title:             "Title",
		Price:             1000,
		ImageUrl:          "https://image.com",
		Author:            "Author",
		Genre:             "Genre",
		YearOfPublication: 2021,
	}
	s.bookRepo.On("FindByID", context.Background(), int64(1)).Return(&domain.Book{}, nil)
	s.bookRepo.On("FindByTitleAndAuthor", context.Background(), req.Title, req.Author).
		Return(&domain.Book{}, nil)
	err := s.useCase.Edit(context.Background(), req, int64(1))
	s.Error(err)
	s.Equal(domain.ErrBookAlreadyExist, err)
}

// Success
func (s *BookSuite) TestEdit() {
	req := domain.EditBookRequest{
		Title:             "Title",
		Price:             1000,
		ImageUrl:          "https://image.com",
		Author:            "Author",
		Genre:             "Genre",
		YearOfPublication: 2021,
	}
	s.bookRepo.On("FindByID", context.Background(), int64(1)).Return(&domain.Book{}, nil)
	s.bookRepo.On("FindByTitleAndAuthor", context.Background(), req.Title, req.Author).
		Return(nil, nil)
	s.bookRepo.On("Update", context.Background(), int64(1), mock.AnythingOfType("*domain.Book")).Return(nil)
	err := s.useCase.Edit(context.Background(), req, int64(1))
	s.NoError(err)
}
