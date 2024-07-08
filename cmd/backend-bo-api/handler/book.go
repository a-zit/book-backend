package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"book-backend/domain"
)

type bookHandler struct {
	bookUseCase domain.BookUseCase
}

func InitBookHandler(g *echo.Group, bookUseCase domain.BookUseCase) {
	b := &bookHandler{
		bookUseCase,
	}
	book := g.Group("/books")
	book.GET("/", b.GetAll)
	book.GET("/:id", b.GetById)
	book.POST("/", b.Create)
	book.PUT("/:id", b.Update)
	book.DELETE("/:id", b.Delete)
}

func (b *bookHandler) GetAll(c echo.Context) error {
	books, err := b.bookUseCase.GetAll(c.Request().Context())
	if err != nil {
		return domain.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, books)
}

func (b *bookHandler) GetById(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return domain.ErrInvalidInteger
	}
	book, err := b.bookUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (b *bookHandler) Create(c echo.Context) error {
	var request domain.CreateBookRequest
	err := c.Bind(&request)
	if err != nil {
		return domain.ErrBadRequest
	}
	err = b.bookUseCase.Create(c.Request().Context(), request)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func (b *bookHandler) Update(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return domain.ErrInvalidInteger
	}

	var request domain.EditBookRequest
	err = c.Bind(&request)
	if err != nil {
		return domain.ErrBadRequest
	}

	err = b.bookUseCase.Edit(c.Request().Context(), request, id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func (b *bookHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return domain.ErrInvalidInteger
	}
	err = b.bookUseCase.Delete(c.Request().Context(), id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
