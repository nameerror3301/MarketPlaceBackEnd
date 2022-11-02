package routes

import (
	"products_go/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RespStatus(apiVersion string, statusCode int, description string) fiber.Map {
	return fiber.Map{
		"api_version": apiVersion,
		"status_code": statusCode,
		"description": description,
	}
}

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Hello, it is starting page"))
}

func GetAll(c *fiber.Ctx) error {
	requ, err := models.ReadCSV(c.Query("total")) // Test
	if requ == nil || err != nil {
		logrus.Warnf("The user transmitted a query value that is not in the valid range (0 or >= 10000) - [%v] [%s]\n", requ, err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(RespStatus("1.0", fiber.ErrBadRequest.Code, "Incorrect data"))
	}

	return c.Status(fiber.StatusOK).Send(requ)
}

func GetById(c *fiber.Ctx) error {
	requ, err := models.GetByIdProduct(c.Params("productId"))
	if requ == nil || err != nil {
		logrus.Warnf("The user transmitted a query value that is not in the valid range (0 or >= 10000) - [%v] [%s]\n", requ, err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(RespStatus("1.0", fiber.ErrBadRequest.Code, "Incorrect data"))
	}

	return c.Status(fiber.StatusOK).Send(requ)
}
