package routes

import (
	models "MarketPlaceBackEnd/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Struct for responce
func RespStatus(apiVersion string, statusCode int, description string) fiber.Map {
	return fiber.Map{
		"api_version": apiVersion,
		"status_code": statusCode,
		"description": description,
	}
}

// WORK: Home page
func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Hello, it is starting page"))
}

// WORK: Get all product and total product (1 - 60)
func GetAll(c *fiber.Ctx) error {
	resp, err := models.ReadCSV(c.Query("total")) // Test
	if resp == nil || err != nil {
		logrus.Warnf("Incorrect data or err func - [%v] [%s]\n", resp, err)
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data"))
	}

	return c.Status(fiber.StatusOK).Send(resp)
}

// WORK: Get product by id
func GetById(c *fiber.Ctx) error {
	resp, err := models.GetByIdProduct(c.Params("productId"))
	if resp == nil || err != nil {
		logrus.Warnf("Incorrect data or err func - [%v] [%s]\n", resp, err)
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data"))
	}

	return c.Status(fiber.StatusOK).Send(resp)
}

// WORK: Create user
func SignUp(c *fiber.Ctx) error {
	if content := c.Request().Header.ContentType(); string(content) != "application/json" {
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "Incorrect Content-type"))
	}
	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Cool!"))
}
