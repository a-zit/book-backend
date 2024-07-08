package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"book-backend/cmd/backend-bo-api/config"
	"book-backend/cmd/backend-bo-api/handler"
	mw "book-backend/cmd/backend-bo-api/middleware"
	"book-backend/database"
	"book-backend/module/book"
	"book-backend/pkg/validator"
)

func main() {
	_ = context.Background()
	e := echo.New()

	// Configure CORS middleware options
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/health", healthCheck)
	g := e.Group("/api")
	c := config.InitConfig()

	// maria db
	mrDB := database.NewMariaDB(c.MariaURI)
	defer mrDB.Close()
	// struct validator
	validator.InitValidator()
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = mw.EchoErrorHandler
	// repo
	bookRepo := book.NewBookRepository(mrDB)
	// useCase
	bookUseCase := book.NewBookUseCase(bookRepo)
	// handler
	handler.InitBookHandler(g, bookUseCase)

	if c.BOPort == "" {
		c.BOPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + c.BOPort))
}

func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Well Done!")
}
