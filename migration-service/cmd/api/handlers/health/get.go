package health

import "github.com/gofiber/fiber/v2"

// health representa la respuesta de estado de salud de la aplicación.
// @Description Detalles del estado de salud de la aplicación.
// @Model health
type health struct {
	Status string `json:"status" example:"ok"`
}

// Check if the service is up
//
// @Summary        Health Check
// @Description    Check if the service is up
// @Tags           Health
// @Produce        json
// @Success        200  {object}  health  "The service is up"
// @Router         /health [get]
func (h Handler) Check(c *fiber.Ctx) error {
	return c.Status(200).JSON(health{Status: "ok"})
}
