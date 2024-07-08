package book_test

import (
	"book-backend/module/book/book_suite"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(book_suite.BookSuite))
}
