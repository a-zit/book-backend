package domain

import appErr "book-backend/pkg/error"

var (
	ErrInternalServerError appErr.Error = appErr.NewError(
		appErr.INTERNAL_SERVER_ERROR,
		"internal-server-error",
		"",
	)
	ErrForbidden        appErr.Error = appErr.NewError(appErr.FORBIDDEN, "forbidden", "")
	ErrBadRequest       appErr.Error = appErr.NewError(appErr.BAD_REQUEST, "bad-request", "")
	ErrValidationFailed appErr.Error = appErr.NewError(appErr.BAD_REQUEST, "validation-failed", "")
	ErrBindStructFailed appErr.Error = appErr.NewError(appErr.BAD_REQUEST, "bind-struct-failed", "")
	ErrInvalidInteger   appErr.Error = appErr.NewError(appErr.BAD_REQUEST, "invalid-integer", "")
)

// book
var (
	ErrBookNotFound     appErr.Error = appErr.NewError(appErr.NOT_FOUND, "book-not-found", "")
	ErrBookAlreadyExist appErr.Error = appErr.NewError(appErr.CONFLICT, "book-already-exist", "")
)

type HTTPErrResponse struct {
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
}
