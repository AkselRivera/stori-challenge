package core

import (
	"errors"

	"github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

func RespondError(c *fiber.Ctx, err error) error {

	var customError domain.CustomError

	if errors.As(err, &customError) {
		if status, ok := ErrorMapper[customError.Code]; ok {
			return c.Status(status).JSON(customError)

		}
	}

	return c.Status(fiber.StatusInternalServerError).JSON(domain.CustomError{
		Code:    domain.ErrorCodeInternalServerError,
		Message: "Something went wrong, please try it again later",
	})
}

var ErrorMapper map[string]int = map[string]int{
	domain.ErrorCodeBadRequest:          fiber.ErrBadRequest.Code,
	domain.ErrorCodeConflict:            fiber.ErrBadRequest.Code,
	domain.ErrorCodeInternalServerError: fiber.ErrInternalServerError.Code,
}
