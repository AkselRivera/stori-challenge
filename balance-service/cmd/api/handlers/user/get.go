package user

import (
	"fmt"
	"time"

	"github.com/AkselRivera/stori-challenge/balance-service/cmd/api/core"
	"github.com/AkselRivera/stori-challenge/balance-service/internal/domain"
	"github.com/gofiber/fiber/v2"
)

// GetBalance user account details
//
// @Summary        Get user balance
// @Description    Retrieve the balance of a user with optional date range.
// @Tags           Balance
// @Accept         json
// @Produce        json
// @Param          user_id  path     int     true  "User ID"
// @Param          from     query    string  false "Start date in RFC3339 format" example(2006-01-02T15:04:05Z07:00)
// @Param          to       query    string  false "End date in RFC3339 format" example(2006-01-02T15:04:05Z07:00)
// @Success        200      {object}  domain.UserBalance  "Successful response with user balance"
// @Failure        400      {object}  domain.CustomError      "Error response for invalid input"
// @Router         /user/{user_id}/balance [get]
func (h Handler) GetBalance(c *fiber.Ctx) error {

	id, err := c.ParamsInt("user_id")

	if err != nil {
		return core.RespondError(c, domain.HandleError(domain.ErrorIdRequired, "numeric value expected for 'user_id'"))
	}

	queryStartDate := c.Query("from")
	queryEndDate := c.Query("to")

	var startDate time.Time
	var endDate time.Time

	if queryStartDate != "" || queryEndDate != "" {
		startDate, err = time.Parse(time.RFC3339, queryStartDate)
		if err != nil {
			return core.RespondError(c, domain.HandleError(domain.ErrorInvalidDate, fmt.Sprintf("%s expected: %s", queryStartDate, time.RFC3339)))
		}

		endDate, err = time.Parse(time.RFC3339, queryEndDate)
		if err != nil {
			return core.RespondError(c, domain.HandleError(domain.ErrorInvalidDate, fmt.Sprintf("%s expected: %s", queryEndDate, time.RFC3339)))
		}
	}

	userBalance, err := h.UserService.GetBalance(id, startDate, endDate)

	if err != nil {
		return core.RespondError(c, err)
	}

	return c.Status(200).JSON(userBalance)
}
