package routes

import (
	"MarketPlaceBackEnd/CodeBack/internal/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"description": "Api v1.0",
		"status":      true,
	})
}

func AllProduct(c *fiber.Ctx) error {
	var p models.ProductData
	request, err := p.ReadCSV("./internal/models/csv/yandex.csv")
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).Send(request)
}
