package main

import (
	"MarketPlaceBackEnd/internal/app"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := app.Run(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Server up!!")
}
