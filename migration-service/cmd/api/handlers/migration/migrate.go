package migration

import (
	"encoding/csv"

	"github.com/AkselRivera/stori-challenge/migration-service/cmd/api/core"
	domain "github.com/AkselRivera/stori-challenge/migration-service/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

// Migrate csv file to database
//
//	@Summary        Migrate csv file to database
//	@Description    Migrate csv file to database
//	@Tags            Migrate
//	@Accept            multipart/form-data
//
// @Param          file  formData  file  true  "CSV file to upload"
//
//	@Produce            application/json
//	@Success        200 {string} string "processing csv file in the background, you will receive an email when the process is completed"
//
// @Failure        400    {object}  domain.CustomError  "Bad request error"
// @Failure        409    {object}  domain.CustomError  "Conflict error"
// @Failure        500    {object}  domain.CustomError  "Internal server error"
// @Router            /migrate [post]
func (h Handler) Migrate(c *fiber.Ctx) error {
	uploadFile, err := c.FormFile("file")

	if err != nil {
		return core.RespondError(c, domain.HandleError(domain.ErrorMissingField, "'file' is required"))
	}

	if uploadFile.Header["Content-Type"][0] != "text/csv" {
		return core.RespondError(c, domain.HandleError(domain.ErrorInvalidFileType, "only 'text/csv' files are allowed"))
	}

	file, err := uploadFile.Open()

	if err != nil {
		return core.RespondError(c, domain.HandleError(domain.ErrorConflict, "could not open file"))
	}

	rows, err := csv.NewReader(file).ReadAll()

	if err != nil {
		return core.RespondError(c, domain.HandleError(domain.ErrorConflict, "could not read file"))
	}

	if err := h.MigrationService.ValidateFileHeaders(rows[0]); err != nil {
		return core.RespondError(c, err)
	}

	transactions, err := h.MigrationService.ValidateData(rows[1:])

	if err != nil {
		return core.RespondError(c, err)
	}

	go h.MigrationService.Migrate(transactions, uploadFile.Filename)

	return c.Status(fiber.StatusOK).JSON("processing csv file in the background, you will receive an email when the process is completed")
}
