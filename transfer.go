package openbank

import (
	"errors"
	"fmt"

	"net/http"

	"github.com/stone-co/go-stone-openbank/types"
)

// TransferService handles communication with Stone Openbank API
type TransferService struct {
	client *Client
}

// DryRunTransfer simulate an Internal or External Transfer
func (s *TransferService) DryRunTransfer(input types.TransferInput, idempotencyKey string) (*types.Transfer, *Response, error) {
	path := "/api/v1/dry_run"
	return s.transfer(input, idempotencyKey, path)
}

// Transfer makes Internal or External Transfer
func (s *TransferService) Transfer(input types.TransferInput, idempotencyKey string) (*types.Transfer, *Response, error) {
	path := "/api/v1"
	return s.transfer(input, idempotencyKey, path)
}

func (s *TransferService) transfer(input types.TransferInput, idempotencyKey, path string) (*types.Transfer, *Response, error) {
	var externalTransfer bool

	if input.Amount == 0 {
		return nil, nil, errors.New("amount can't be 0")
	}
	if input.AccountID == "" {
		return nil, nil, errors.New("account_id can't be empty")
	}
	if input.Target.Account.AccountCode == "" {
		return nil, nil, errors.New("account_code can't be empty")
	}

	if input.Target.Account.InstitutionCode != "" {
		if input.Target.Account.BranchCode == "" {
			return nil, nil, errors.New("branch_code can't be empty")
		}
		if input.Target.Entity.Name == "" {
			return nil, nil, errors.New("entity name can't be empty")
		}
		if input.Target.Entity.Document == "" {
			return nil, nil, errors.New("entity document can't be empty")
		}
		if input.Target.Entity.DocumentType == "" {
			return nil, nil, errors.New("entity document type can't be empty")
		}
		externalTransfer = true
	}

	if externalTransfer {
		path = fmt.Sprintf("%s/external_transfers", path)
	} else {
		path = fmt.Sprintf("%s/internal_transfers", path)
	}

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.AddIdempotencyHeader(req, idempotencyKey)
	if err != nil {
		return nil, nil, err
	}

	var transfer types.Transfer
	resp, err := s.client.Do(req, &transfer)
	if err != nil {
		return nil, resp, err
	}

	return &transfer, resp, err
}

// ListInternal returns a list of internal_transfers
func (s *TransferService) ListInternal(accountID string) ([]types.Transfer, *Response, error) {
	path := fmt.Sprintf("/api/v1/internal_transfers?account_id=%s", accountID)
	return s.list(path)
}

// ListExternal returns a list of external_transfers
func (s *TransferService) ListExternal(accountID string) ([]types.Transfer, *Response, error) {
	path := fmt.Sprintf("/api/v1/external_transfers?account_id=%s", accountID)
	return s.list(path)
}

func (s *TransferService) list(path string) ([]types.Transfer, *Response, error) {
	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor     `json:"cursor"`
		Data   []types.Transfer `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}

// GetInternal returns an internal transfer
func (s *TransferService) GetInternal(transferID string) (*types.Transfer, *Response, error) {
	path := fmt.Sprintf("/api/v1/internal_transfers/%s", transferID)
	return s.get(path)
}

// GetExternal returns an external transfer
func (s *TransferService) GetExternal(transferID string) (*types.Transfer, *Response, error) {
	path := fmt.Sprintf("/api/v1/external_transfers/%s", transferID)
	return s.get(path)
}

func (s *TransferService) get(path string) (*types.Transfer, *Response, error) {
	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var transfer types.Transfer
	resp, err := s.client.Do(req, &transfer)
	if err != nil {
		return nil, resp, err
	}

	return &transfer, resp, err
}

// CancelInternal cancels a scheduled internal transference
func (s *TransferService) CancelInternal(transferID string) (*Response, error) {
	path := fmt.Sprintf("/api/v1/internal_transfers/%s/cancel", transferID)
	return s.cancel(path)
}

// CancelExternal cancels a scheduled external transference
func (s *TransferService) CancelExternal(transferID string) (*Response, error) {
	path := fmt.Sprintf("/api/v1/external_transfers/%s/cancel", transferID)
	return s.cancel(path)
}

func (s *TransferService) cancel(path string) (*Response, error) {
	req, err := s.client.NewAPIRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
