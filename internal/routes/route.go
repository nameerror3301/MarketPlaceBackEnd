package routes

import (
	"MarketPlaceBackEnd/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RespStatus(apiVersion string, statusCode int, description string) fiber.Map {
	return fiber.Map{
		"apiVersion":  apiVersion,
		"statusCode":  statusCode,
		"description": description,
	}
}

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Hello, it is starting page"))
}

func GetAll(c *fiber.Ctx) error {
	resp, err := models.GetAllProducts()
	if err != nil {
		logrus.Warn(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(RespStatus("1.0", fiber.ErrBadRequest.Code, "Incorrect data"))
	}

	return c.Status(fiber.StatusOK).Send(resp)
}

func GetById(c *fiber.Ctx) error {
	resp, err := models.GetByIdProduct(c.Params("productId"))
	if err != nil || resp == nil {
		logrus.Warn(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(RespStatus("1.0", fiber.ErrBadRequest.Code, "Incorrect data"))
	}

	return c.Status(fiber.StatusOK).Send(resp)
}
