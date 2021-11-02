package openbank

import (
	"fmt"
	"net/http"

	"github.com/stone-co/go-stone-openbank/types"
)

type InstitutionService struct {
	client *Client
}

type InstitutionContext string

const (
	AllInstitutions InstitutionContext = "all"
	SPIParticipants InstitutionContext = "spi"
	STRParticipants InstitutionContext = "str"
)

// Get institution info
func (s InstitutionService) Get(context string) (*types.Institution, *Response, error) {

	path := fmt.Sprintf("/api/v1/institutions/%s", context)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, emptyIdempotencyKey, nil)
	if err != nil {
		return nil, nil, err
	}

	var institution types.Institution
	resp, err := s.client.Do(req, &institution)
	if err != nil {
		return nil, resp, err
	}

	return &institution, resp, err
}

// List institutions
func (s InstitutionService) List(context InstitutionContext) ([]types.Institution, *Response, error) {

	path := fmt.Sprintf("/api/v1/institutions?context=%s", context)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, emptyIdempotencyKey, nil)
	if err != nil {
		return nil, nil, err
	}

	var institution []types.Institution
	resp, err := s.client.Do(req, &institution)
	if err != nil {
		return nil, resp, err
	}

	return institution, resp, err
}
