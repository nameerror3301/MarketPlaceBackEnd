package routes

import (
	"time"

	status "MarketPlaceBackEnd/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

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

// WORK: Create user
func SignUp(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	if content := c.Request().Header.ContentType(); string(content) != "application/json" {
		return c.Status(fiber.StatusBadRequest).JSON(status.RespStatus("1.0", fiber.StatusBadRequest, "Incorrect Content-type", nil))
	}

	token, err := GenJwtToket()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(status.RespStatus("1.0", fiber.StatusServiceUnavailable, "Technical failures", nil))
	}
	return c.Status(fiber.StatusOK).JSON(status.RespStatus("1.0", fiber.StatusOK, token, nil))
}

func SignIn(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	if content := c.Request().Header.ContentType(); string(content) != "application/json" {
		return c.Status(fiber.StatusBadRequest).JSON(status.RespStatus("1.0", fiber.StatusBadRequest, "Incorrect Content-type", nil))
	}
	return c.Status(fiber.StatusOK).JSON(status.RespStatus("1.0", fiber.StatusOK, "It is sign-in!", nil))
}
