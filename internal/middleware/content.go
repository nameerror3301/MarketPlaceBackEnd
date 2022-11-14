package middleware

import (
	responce "MarketPlaceBackEnd/internal/handler"

	"github.com/gofiber/fiber/v2"
)

/*
if content := c.Request().Header.ContentType(); string(content) != "application/json" {
			return c.Status(fiber.StatusBadRequest).JSON(responce.RespStatus("1.0", fiber.StatusBadRequest, "Incorrect Content-Type", nil))
		}
*/

// WORK
func CheckContentType() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if content := c.Request().Header.ContentType(); string(content) != "application/json" {
			return c.Status(fiber.StatusBadRequest).JSON(responce.RespStatus("1.0", fiber.StatusBadRequest, "Incorrect Content-Type", nil))
		}
		return c.Next()
	}
}
