package routes

import (
	status "MarketPlaceBackEnd/internal/handler"
	models "MarketPlaceBackEnd/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// WORK: Get all product and total product
func GetAll(c *fiber.Ctx) error {
	resp, err := models.ReadCSV(c.Query("total")) // Test
	if resp == nil || err != nil {
		logrus.Warnf("Incorrect data or err func - [%v] [%s]\n", resp, err)
		return c.Status(fiber.StatusBadRequest).JSON(status.RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data", resp))
	}

	return c.Status(fiber.StatusOK).JSON(status.RespStatus("1.0", fiber.StatusOK, "All products", resp))
}

// WORK: Get product by id
func GetById(c *fiber.Ctx) error {
	resp, err := models.GetByIdProduct(c.Params("productId"))
	if resp == nil || err != nil {
		logrus.Warnf("Incorrect data or err func - [%v] [%s]\n", resp, err)
		return c.Status(fiber.StatusBadRequest).JSON(status.RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data", nil))
	}

	return c.Status(fiber.StatusOK).JSON(status.RespStatus("1.0", fiber.StatusOK, "By id product", resp))
}
