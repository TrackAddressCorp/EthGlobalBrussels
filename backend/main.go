package main

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
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
	db.AddPetition(models.Petition{Title: "Test", Description: "Test"})
}

func initFiber() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
