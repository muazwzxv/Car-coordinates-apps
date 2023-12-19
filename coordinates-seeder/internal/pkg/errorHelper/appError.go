package errorHelper

import (
	"bytes"
	"fmt"
)

var (
	ErrMissingName ErrorDetail = ErrorDetail{
		Message: "NAME IS MISSING",
		Code:    "MISSING_NAME",
	}

	ErrMissingBrand ErrorDetail = ErrorDetail{
		Message: "BRAND IS MISSING",
		Code:    "MISSING_BRAND",
	}

	ErrMissingType ErrorDetail = ErrorDetail{
		Message: "TYPE IS MISSING",
		Code:    "MISSING_TYPE",
	}

	ErrInvalidType ErrorDetail = ErrorDetail{
		Message: "TYPE IS INVALID",
		Code:    "INVALID_TYPE",
	}

	ErrMissingBuildDate ErrorDetail = ErrorDetail{
		Message: "BUILD DATE IS MISSING",
		Code:    "MISSING_BUILD_DATE",
	}
)

type ErrorDetail struct {
	Message string
	Code    string
}

type ErrorDetails []ErrorDetail

func (e ErrorDetail) Error() string {
	return fmt.Sprintf("Message: %s \n Code: %s \n\n", e.Message, e.Code)
}

func (errs ErrorDetails) Error() string {
	var errBuf bytes.Buffer

	for _, err := range errs {
    errBuf.WriteString(err.Error())
	}

  return errBuf.String()
}
