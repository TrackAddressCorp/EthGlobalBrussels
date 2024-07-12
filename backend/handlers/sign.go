package handlers

import (
	"github.com/TrackAddressCorp/EthGlobalBrussels/db"
	"github.com/TrackAddressCorp/EthGlobalBrussels/worldid"
	"github.com/gofiber/fiber/v2"
)

type SignRequest struct {
	NullifierHash     string `json:"nullifier_hash"`
	Proof             string `json:"proof"`
	MerkleRoot        string `json:"merkle_root"`
	VerificationLevel string `json:"verification_level"`
	Action            string `json:"action"`
	SignalHash        string `json:"signal_hash"`

	PetitionID uint `json:"petition_id"`
}

func SignPetition(c *fiber.Ctx) error {
	var signRequest SignRequest
	if err := c.BodyParser(&signRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	successResponse, err := worldid.VerifyWorldIDProof(&worldid.VerifyRequest{
		NullifierHash:     signRequest.NullifierHash,
		Proof:             signRequest.Proof,
		MerkleRoot:        signRequest.MerkleRoot,
		VerificationLevel: signRequest.VerificationLevel,
		Action:            signRequest.Action,
		SignalHash:        signRequest.SignalHash,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := db.AddSign(signRequest.PetitionID, successResponse.NullifierHash); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to sign")
	}

	return c.SendString("Signed successfully")
}
