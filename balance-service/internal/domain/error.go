package domain

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
)

const (
	ErrorCodeNotFound            = "not_found"
	ErrorCodeBadRequest          = "bad_request"
	ErrorCodeInternalServerError = "internal_server_error"
)

// Centinel Errors
var ErrorUserNotFound = errors.New("user not found")
var ErrorIdRequired = errors.New("id is required")
var ErrorInvalidDate = errors.New("invalid date format")

// CustomError
// @Description An error that includes a specific code and a message with more details.
// @Model CustomError
type CustomError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)

}

func HandleError(err error, message string) error {
	var customError CustomError

	switch {
	case errors.Is(err, ErrorUserNotFound):
		log.Error("user not found")
		customError = CustomError{
			Code:    ErrorCodeBadRequest,
			Message: "user not found",
		}

	case errors.Is(err, ErrorIdRequired):
		log.Error("id is required")
		customError = CustomError{
			Code:    ErrorCodeBadRequest,
			Message: "id is required",
		}

	case errors.Is(err, ErrorInvalidDate):
		log.Error("invalid date format")
		customError = CustomError{
			Code:    ErrorCodeBadRequest,
			Message: "invalid date format",
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
