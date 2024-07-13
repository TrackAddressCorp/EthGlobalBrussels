package handlers

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"github.com/gofiber/fiber/v2"
)

func ListPetitions(c *fiber.Ctx) error {
	petitions, err := db.ListPetitions()
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Response{StatusMsg: err.Error(), StatusCode: fiber.StatusForbidden})
	}

	return c.JSON(petitions)
}
