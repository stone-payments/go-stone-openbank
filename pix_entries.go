package openbank

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/stone-co/go-stone-openbank/types"
)

//ListEntries list the PIX keys of an account
func (s *PixService) ListEntries(accountID string) ([]types.PixEntry, *Response, error) {
	path := fmt.Sprintf("/api/v1/pix/%s/entries", accountID)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor     `json:"cursor"`
		Data   []types.PixEntry `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}

// CreateEntry creates a new Key Entry
func (s *PixService) CreateEntry(input types.CreateEntryInput, idempotencyKey string) (string, *Response, error) {
	if input.AccountID == "" {
		return "", nil, errors.New("accountID cannot be empty")
	}

	path := fmt.Sprintf("/api/v1/pix/%s/entries", input.AccountID)

	if input.ParticipantISPB == "" {
		input.ParticipantISPB = "16501555"
	}

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return "", nil, err
	}

	err = s.client.AddIdempotencyHeader(req, idempotencyKey)
	if err != nil {
		return "", nil, err
	}

	if input.VerificationID != "" {
		req.Header.Add("x-stone-verification-id", input.VerificationID)
	}
	if input.VerificationCode != "" {
		req.Header.Add("x-stone-verification-code", input.VerificationCode)
	}

	var output struct{ ID string }
	resp, err := s.client.Do(req, &output)
	if err != nil {
		return "", resp, err
	}

	return output.ID, resp, err
}
