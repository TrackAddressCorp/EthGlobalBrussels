package handlers

import (
	"strconv"

	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"github.com/gofiber/fiber/v2"
)

func FinishPetition(c *fiber.Ctx) error {
	id := c.Params("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{StatusMsg: "Invalid ID", StatusCode: fiber.StatusBadRequest})
	}

	err = db.SetPetitionFinished(uint(intID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{StatusMsg: "Petition not found", StatusCode: fiber.StatusNotFound})
	}

	return c.SendString("Finish Petition")
}
