package app

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

func Run() error {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// fmt.Errorf("Err up server - %s", err)
	if err := app.Listen(":8080"); err != nil {
		return fmt.Errorf("Err up server - %s", err)
	}
	return nil
}
