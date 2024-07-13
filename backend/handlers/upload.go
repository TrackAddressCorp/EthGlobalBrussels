package handlers

import (
	"os"
	"strconv"

	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/lighthouse"
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"github.com/gofiber/fiber/v2"
)

type uploadResponse struct {
	models.Response
	PdfURL string `json:"pdf_url"`
}

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("pdf")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{StatusMsg: "Invalid request, file is missing", StatusCode: fiber.StatusBadRequest})
	}

	petitionIDString := c.FormValue("petition_id")
	if petitionIDString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{StatusMsg: "Invalid request", StatusCode: fiber.StatusBadRequest})
	}
	petitionID, err := strconv.Atoi(petitionIDString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{StatusMsg: "Invalid petition ID", StatusCode: fiber.StatusBadRequest})
	}

	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		os.Mkdir("./uploads", 0755)
	}

	err = c.SaveFile(file, "./uploads/"+file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{StatusMsg: "Failed to save file", StatusCode: fiber.StatusInternalServerError})
	}

	resp, err := lighthouse.UploadFile("./uploads/" + file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{StatusMsg: "Failed to upload file to lighthouse", StatusCode: fiber.StatusInternalServerError})
	}

	err = db.SetPdfURL(uint(petitionID), "https://gateway.lighthouse.storage/ipfs/"+resp.Hash)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{StatusMsg: "Failed to save file URL", StatusCode: fiber.StatusInternalServerError})
	}

	return c.JSON(uploadResponse{Response: models.Response{StatusMsg: "File uploaded successfully", StatusCode: fiber.StatusOK}, PdfURL: "https://gateway.lighthouse.storage/ipfs/" + resp.Hash})
}
