package main

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	initDB()
	initFiber()
}

func initDB() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
}

func initFiber() {
	app := fiber.New()

	app.Get("/petition/:id", handlers.GetPetition)
	app.Get("/petitions", handlers.ListPetitions)
	app.Post("/petition/sign", handlers.SignPetition)
	app.Post("/petition/create", handlers.CreatePetition)

	app.Listen(":4242")
}
