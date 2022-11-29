package app

import (
	middle "MarketPlaceBackEnd/internal/middleware"
	"MarketPlaceBackEnd/internal/routes"

	conn "MarketPlaceBackEnd/internal/database"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

// middle.CheckContentType()
func setUpRoutes(app *fiber.App) {
	app.Get("/api/v1", routes.Home)
	app.Get("/api/v1/products", middle.CheckJwtToken(), routes.GetAll)
	app.Get("/api/v1/products/:productId", middle.CheckJwtToken(), routes.GetById)
	app.Post("/api/v1/signUp", middle.CheckContentType(), routes.SignUp)
	app.Post("/api/v1/signIn", middle.CheckContentType(), routes.SignIn)
}

func Run() {
	app := fiber.New()
	app.Use(logger.New())

	setUpRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		logrus.Fatalf("Err up server - %s", err)
	}

	if db, err := conn.ConnectDatabase(); db == nil || err != nil {
		logrus.Warnf("Err connect to database - %s", err)
	} else {
		logrus.Info("Connection to the database was successful")
		db.Close()
	}

	logrus.Info("Service is up!")
}
