package handlers

import (
	"strconv"

	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/gofiber/fiber/v2"
)

func GetPartition(c *fiber.Ctx) error {
	id := c.Params("id")

	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	partition, err := db.GetPetition(uint(intID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Partition not found")
	}

	return c.JSON(partition)
}
