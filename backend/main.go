package main

import (
	"log"

	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/handlers"
	"github.com/TrackAddressCorp/EthGlobalBrussels/worldid"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	actionresponse, errorresponse, err := worldid.CreateDynamicAction(worldid.CreateActionRequest{
		Action: "petition", MaxVerifications: 1,
	})
	log.Println(actionresponse)
	log.Println(errorresponse)
	log.Println(err)

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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/petition/:id", handlers.GetPetition)
	app.Get("/petitions", handlers.ListPetitions)
	app.Post("/petition/sign", handlers.SignPetition)
	app.Post("/petition/create", handlers.CreatePetition)

	app.Listen(":4242")
}
