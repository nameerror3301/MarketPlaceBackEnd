package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func CheckJwtToken() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   []byte("SECRET"),
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"api_version": "1.0", "status_code": fiber.StatusUnauthorized, "description": "Incorrect or invalid access token!"})
}
