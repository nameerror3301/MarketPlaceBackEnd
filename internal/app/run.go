package app

import (
	"MarketPlaceBackEnd/internal/routes"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/api/v1", routes.Home)
	app.Get("/api/v1/all", routes.AllProduct)

}

func Run() error {
	app := fiber.New()
	app.Use(logger.New())

	SetUpRoutes(app)

	// fmt.Errorf("Err up server - %s", err)
	if err := app.Listen(":8080"); err != nil {
		return fmt.Errorf("Err up server - %s", err)
	}
	return nil
}
