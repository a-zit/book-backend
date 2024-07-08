package app_error

import (
	"fmt"
	"net/http"
)

const (
	BAD_REQUEST           string = "bad-request"
	FORBIDDEN             string = "forbidden"
	NOT_FOUND             string = "not-found"
	CONFLICT              string = "conflict"
	INTERNAL_SERVER_ERROR string = "internal-server-error"
)

type Error struct {
	Category string
	Code     string
	Message  interface{}
}

func NewError(category, code string, message interface{}) Error {
	return Error{
		Category: category,
		Code:     code,
		Message:  message,
	}
}

func (e Error) Error() string {
	m := e.Code
	if e.Message != nil && e.Message != "" {
		m = fmt.Sprintf("%s: %v", m, e.Message)
	}
	return m
}

func (e Error) SetMessage(message interface{}) Error {
	e.Message = message
	return e
}

func (e Error) GetHTTPStatusCode() int {
	switch e.Category {
	case BAD_REQUEST:
		return http.StatusBadRequest
	case FORBIDDEN:
		return http.StatusForbidden
	case NOT_FOUND:
		return http.StatusNotFound
	case CONFLICT:
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}
