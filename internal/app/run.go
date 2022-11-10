package app

import (
	middle "MarketPlaceBackEnd/internal/middleware"
	routes "MarketPlaceBackEnd/internal/routes"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/api/v1", routes.Home)
	app.Get("/api/v1/products", middle.CheckJwtToken(), routes.GetAll)
	app.Get("/api/v1/products/:productId", middle.CheckJwtToken(), routes.GetById)
	app.Post("/api/v1/signUp", routes.SignUp)
}

func Run() {
	app := fiber.New()
	app.Use(logger.New())

	SetUpRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		logrus.Fatalf("Err up server - %s", err)
	}

	logrus.Info("Service is up!")
}