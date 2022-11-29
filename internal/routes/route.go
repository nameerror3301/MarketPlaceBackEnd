package routes

import (
	models "MarketPlaceBackEnd/internal/models"

	"github.com/gofiber/fiber/v2"
)

// Struct for responce
func RespStatus(apiVersion string, statusCode int, description string, content []models.ProductData) fiber.Map {
	return fiber.Map{
		"api_version": apiVersion,
		"status_code": statusCode,
		"description": description,
		"content":     content,
	}
}

// WORK: Home page
func Home(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Hello, it is starting page", nil))
}
