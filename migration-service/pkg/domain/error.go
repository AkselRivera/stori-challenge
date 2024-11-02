package domain

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
)

const (
	ErrorCodeBadRequest          = "bad_request"
	ErrorCodeConflict            = "conflict"
	ErrorCodeInternalServerError = "internal_server_error"
)

var ErrorInvalidCsvColumns = errors.New("invalid csv columns")
var ErrorInvalidFileType = errors.New("invalid file type")
var ErrorInvalidDataType = errors.New("invalid data type")
var ErrorMissingField = errors.New("missing field")
var ErrorConflict = errors.New("conflict")

// CustomError
// @Description An error that includes a specific code and a message with more details.
// @Model CustomError
type CustomError struct {
	Code    string `json:"code" example:"error_code"`
	Message string `json:"message" example:"error message"`
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func HandleError(err error, message string) error {
	var customError CustomError

	switch {
	case errors.Is(err, ErrorInvalidCsvColumns):
		log.Error("invalid csv columns")
		customError = CustomError{
			Code:    ErrorCodeBadRequest,
			Message: "invalid csv columns",
		}

	case errors.Is(err, ErrorInvalidFileType):
		log.Error("invalid file type")
		customError = CustomError{
			Code:    ErrorCodeBadRequest,
			Message: "invalid file type",
		}

	case errors.Is(err, ErrorInvalidDataType):
		log.Error("invalid data type")
		customError = CustomError{
			Code:    ErrorCodeBadRequest,
			Message: "invalid data type",
		}

	case errors.Is(err, ErrorMissingField):
		log.Error("missing field")
		customError = CustomError{
			Code:    ErrorCodeBadRequest,
			Message: "missing field",
		}

	case errors.Is(err, ErrorConflict):
		log.Error("missing field")
		customError = CustomError{
			Code:    ErrorCodeConflict,
			Message: "conflict",
		}

	default:
		log.Error(err.Error())
		customError = CustomError{
			Code:    ErrorCodeInternalServerError,
			Message: "something went wrong, please try it again later",
		}
	}

	if message != "" && customError.Code != ErrorCodeInternalServerError {
		customError.Message = fmt.Sprintf("%s: %s", customError.Message, message)
	}

	return customError
}
