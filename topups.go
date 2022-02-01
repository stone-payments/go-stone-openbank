package openbank

import (
	"fmt"
	"net/http"

	"github.com/stone-co/go-stone-openbank/types"
)

// TopupsService handles communication with Stone Openbank API
type TopupsService struct {
	client *Client
}

// ListGameProviders list all game providers
func (s *TopupsService) ListGameProviders() (*types.Providers, *Response, error) {
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

// GetValuesFromGameProvider list all values from a game provider
func (s *TopupsService) GetValuesFromGameProvider(id int) (*types.Products, *Response, error) {
	path := fmt.Sprintf("/api/v1/topups/games/values/%v", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var products types.Products
	resp, err := s.client.Do(req, &products)
	if err != nil {
		return nil, resp, err
	}

	return &products, resp, err
}
