package main

import (
	"MarketPlaceBackEnd/CodeBack/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
