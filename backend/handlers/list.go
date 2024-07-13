package handlers

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"github.com/gofiber/fiber/v2"
)

type listPetitionsResponse struct {
	models.Response
	Petitions []models.Petition `json:"petitions"`
}

func ListPetitions(c *fiber.Ctx) error {
	petitions, err := db.ListPetitions()
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(models.Response{StatusMsg: err.Error(), StatusCode: fiber.StatusForbidden})
	}

	return c.JSON(listPetitionsResponse{Response: models.Response{StatusMsg: "Petitions fetched successfully", StatusCode: fiber.StatusOK}, Petitions: petitions})
}
