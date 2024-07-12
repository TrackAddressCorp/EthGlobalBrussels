package handlers

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"github.com/gofiber/fiber/v2"
)

type CreatePetitionRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	PdfURL      *string `json:"pdf_url"`
}

type CreatePetitionResponse struct {
	ID uint `json:"id"`
}

func CreatePetition(c *fiber.Ctx) error {
	var petitionRequest CreatePetitionRequest
	if err := c.BodyParser(&petitionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	newPetition := models.Petition{
		Title:       petitionRequest.Title,
		Description: petitionRequest.Description,
		PdfURL:      petitionRequest.PdfURL,
	}

	if err := db.CreatePetition(&newPetition); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}

	return c.JSON(CreatePetitionResponse{ID: newPetition.ID})
}
