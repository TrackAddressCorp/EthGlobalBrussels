package worldid

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

type VerifyRequest struct {
	NullifierHash     string `json:"nullifier_hash"`
	Proof             string `json:"proof"`
	MerkleRoot        string `json:"merkle_root"`
	VerificationLevel string `json:"verification_level"`
	Action            string `json:"action"`
}

type VerifyResponseSuccess struct {
	Success       bool   `json:"success"`
	Action        string `json:"action"`
	NullifierHash string `json:"nullifier_hash"`
	CreatedAt     string `json:"created_at"`
}

type VerifyResponseError struct {
	Code      string      `json:"code"`
	Detail    string      `json:"detail"`
	Attribute interface{} `json:"attribute"`
}

func VerifyWorldIDProof(verifyRequest *VerifyRequest) (*VerifyResponseSuccess, error) {
	jsonData, err := json.Marshal(verifyRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://developer.worldcoin.org/api/v2/verify/"+os.Getenv("WORLDID_APP_ID"), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		var successResponse VerifyResponseSuccess
		err = json.Unmarshal(body, &successResponse)
		// log.Println(string(body))
		if err != nil {
			return nil, err
		}
		return &successResponse, nil
	} else {
		var errorResponse VerifyResponseError
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(errorResponse.Detail)
	}
}
