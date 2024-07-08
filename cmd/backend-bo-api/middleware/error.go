package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"book-backend/domain"
	appErr "book-backend/pkg/error"
)

func EchoErrorHandler(err error, c echo.Context) {
	var (
		httpCode int         = http.StatusInternalServerError
		errCode  string      = appErr.INTERNAL_SERVER_ERROR
		message  interface{} = ""
	)

	appErr, ok := err.(appErr.Error)
	if ok {
		httpCode = appErr.GetHTTPStatusCode()
		errCode = appErr.Code
		message = appErr.Message
	}

	if err := c.JSON(httpCode, domain.HTTPErrResponse{Code: errCode, Message: message}); err != nil {
		return
	}
}
