package worldid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type CreateActionRequest struct {
	Action           string `json:"action"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	MaxVerifications int    `json:"max_verifications,omitempty"`
}

type CreateActionParams struct {
	AppID string `json:"app_id"`
}

type ErrorResponse struct {
	Code      string `json:"code"`
	Detail    string `json:"detail"`
	Attribute string `json:"attribute,omitempty"`
}

type CreateActionResponse struct {
	Action interface{} `json:"action"`
}

func CreateDynamicAction(body CreateActionRequest) (*CreateActionResponse, *ErrorResponse, error) {
	url := "https://developer.worldcoin.org/api/v2/create-action/" + os.Getenv("WORLDID_APP_ID")

	paramsBytes, err := json.Marshal(CreateActionParams{AppID: os.Getenv("WORLDID_APP_ID")})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal params: %w", err)
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("WORLDID_API_KEY"))
	req.URL.RawQuery = string(paramsBytes)

	// Send the request
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp ErrorResponse
		if err := json.Unmarshal(respBody, &errorResp); err != nil {
			return nil, nil, fmt.Errorf("failed to unmarshal error response: %w", err)
		}
		return nil, &errorResp, nil
	}

	// Unmarshal the successful response
	var createActionResp CreateActionResponse
	if err := json.Unmarshal(respBody, &createActionResp); err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal success response: %w", err)
	}

	return &createActionResp, nil, nil
}
