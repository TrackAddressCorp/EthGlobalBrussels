package main

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/handlers"
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
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
	description := "Test"
	db.AddPetition(models.Petition{Title: "asd", Description: &description})
}

func initFiber() {
	app := fiber.New()

	app.Get("/petition/:id", handlers.GetPetition)
	app.Post("/petition/sign", handlers.SignPetition)

	app.Listen(":3000")
}
