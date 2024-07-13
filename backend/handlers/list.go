package handlers

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/gofiber/fiber/v2"
)

func ListPetitions(c *fiber.Ctx) error {
	petitions, err := db.ListPetitions()
	if err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}

	return c.JSON(petitions)
}
