package openbank

import (
	"net/http"

	"github.com/stone-co/go-stone-openbank/types"
)

// TopUpsService handles communication with Stone Openbank API
type TopUpsService struct {
	client *Client
}

// ListGameProviders list all game providers
func (s *TopUpsService) ListGameProviders() (*types.Providers, *Response, error) {
	const path = "api/v1/topups/games/providers"

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var providers types.Providers
	resp, err := s.client.Do(req, &providers)
	if err != nil {
		return nil, resp, err
	}

	return &providers, resp, err
}