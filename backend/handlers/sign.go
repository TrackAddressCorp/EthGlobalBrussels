package handlers

import (
	"strconv"

	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/models"
	"github.com/TrackAddressCorp/EthGlobalBrussels/worldid"
	"github.com/gofiber/fiber/v2"
)

type signRequest struct {
	NullifierHash     string `json:"nullifier_hash"`
	Proof             string `json:"proof"`
	MerkleRoot        string `json:"merkle_root"`
	VerificationLevel string `json:"verification_level"`
	Action            string `json:"action"`
}

func SignPetition(c *fiber.Ctx) error {
	var signRequest signRequest
	if err := c.BodyParser(&signRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{StatusMsg: "Invalid request", StatusCode: fiber.StatusBadRequest})
	}

	actionID, err := strconv.Atoi(signRequest.Action)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{StatusMsg: "Invalid action ID", StatusCode: fiber.StatusBadRequest})
	}

	successResponse, err := worldid.VerifyWorldIDProof(&worldid.VerifyRequest{
		NullifierHash:     signRequest.NullifierHash,
		Proof:             signRequest.Proof,
		MerkleRoot:        signRequest.MerkleRoot,
		VerificationLevel: signRequest.VerificationLevel,
		Action:            signRequest.Action,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{StatusMsg: err.Error(), StatusCode: fiber.StatusBadRequest})
	}

	if err := db.AddSign(uint(actionID), successResponse.NullifierHash); err != nil {
		return c.Status(fiber.StatusForbidden).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{StatusMsg: "Petition signed successfully", StatusCode: fiber.StatusOK})
}
