package errorHelper

import (
	"errors"
)

var (
	ErrBadRequest       = errors.New("bad request")
	ErrInternalServer   = errors.New("internal server error")
	ErrFileNotSupported = errors.New("file type not supported")
	ErrNotFound         = errors.New("not found")
	ErrConflict         = errors.New("conflict")
	ErrForbidden        = errors.New("forbidden")
	ErrUnAuthorized     = errors.New("unauthorized")
)

func ApplicationError(errs ErrorDetails) map[string]ErrorDetails {
	return map[string]ErrorDetails{
		"error": errs,
	}
}

func SimpleErrorResponse(err error) map[string]any {
	return map[string]any{
		"error": err,
	}
}

func ToResponseBody(data interface{}) map[string]any {
	return map[string]any{
		"data": data,
	}
}
