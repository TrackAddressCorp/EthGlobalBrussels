package handlers

import (
	"strconv"

	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"github.com/gofiber/fiber/v2"
)

type petitionResponse struct {
	models.Response
	Petition models.Petition `json:"petition"`
}

func GetPetition(c *fiber.Ctx) error {
	id := c.Params("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{StatusMsg: "Invalid ID", StatusCode: fiber.StatusBadRequest})
	}

	partition, err := db.GetPetition(uint(intID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{StatusMsg: "Petition not found", StatusCode: fiber.StatusNotFound})
	}

	return c.JSON(petitionResponse{
		Response: models.Response{StatusMsg: "Petition fetched successfully", StatusCode: fiber.StatusOK},
		Petition: partition,
	})
}
