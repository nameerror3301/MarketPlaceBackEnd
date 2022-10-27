package routes

import (
	"MarketPlaceBackEnd/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
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
		logrus.Fatal(err)
	}

	return c.Status(fiber.StatusOK).Send(request)
}
