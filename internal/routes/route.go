package routes

import (
	models "MarketPlaceBackEnd/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
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

func GenJwtToket() (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// The secret key will be stored in an environment variable (In the future...)
	t, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		logrus.Fatalf("Error in the process of generating a token for the user - %s", err)
	}
	return t, nil
}

// WORK: Home page
func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "Hello, it is starting page", nil))
}

// WORK: Get all product and total product (1 - 60)
func GetAll(c *fiber.Ctx) error {
	resp, err := models.ReadCSV(c.Query("total")) // Test
	if resp == nil || err != nil {
		logrus.Warnf("Incorrect data or err func - [%v] [%s]\n", resp, err)
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data", resp))
	}

	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "All products", resp))
}

// WORK: Get product by id
func GetById(c *fiber.Ctx) error {
	resp, err := models.GetByIdProduct(c.Params("productId"))
	if resp == nil || err != nil {
		logrus.Warnf("Incorrect data or err func - [%v] [%s]\n", resp, err)
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "Incorrect data", nil))
	}

	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "By id product", resp))
}

// WORK: Create user
func SignUp(c *fiber.Ctx) error {
	// Check content type
	if content := c.Request().Header.ContentType(); string(content) != "application/json" {
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "Incorrect Content-type", nil))
	}

	token, err := GenJwtToket()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(RespStatus("1.0", fiber.StatusServiceUnavailable, "Technical failures", nil))
	}
	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, token, nil))
}

func SignIn(c *fiber.Ctx) error {
	// Check content type
	if content := c.Request().Header.ContentType(); string(content) != "application/json" {
		return c.Status(fiber.StatusBadRequest).JSON(RespStatus("1.0", fiber.StatusBadRequest, "Incorrect Content-type", nil))
	}
	return c.Status(fiber.StatusOK).JSON(RespStatus("1.0", fiber.StatusOK, "It is sign-in!", nil))
}
