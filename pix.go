package openbank

import (
"fmt"
"net/http"

"github.com/stone-co/go-stone-openbank/types"
)

// PixService handles communication with Stone Openbank API
type PixService struct {
	client *Client
}

// GetOutboundPix is a service used to retrieve information details from a Pix.
func (s *PixService) GetOutboundPix(id string) (*types.PixOutBoundOutput, *Response, error) {

	path := fmt.Sprintf("/api/v1/pix/outbound_pix_payments/%s", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var pix types.PixOutBoundOutput
	resp, err := s.client.Do(req, &pix)
	if err != nil {
		return nil, resp, err
	}

	return &pix, resp, err
}
