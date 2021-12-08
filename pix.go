package openbank

import (
"fmt"
"net/http"

"github.com/stone-co/go-stone-openbank/types"
)

// PIXService handles communication with Stone Openbank API
type PIXService struct {
	client *Client
}

// GetOutboundPix is a service used to retrieve information details from a Pix.
func (s *PIXService) GetOutboundPix(id string) (*types.PIXOutBoundOutput, *Response, error) {

	path := fmt.Sprintf("/api/v1/pix/outbound_pix_payments/%s", id)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var pix types.PIXOutBoundOutput
	resp, err := s.client.Do(req, &pix)
	if err != nil {
		return nil, resp, err
	}

	return &pix, resp, err
}

// GetQRCodeData is a service used to retrieve information details from a Pix QRCode.
func (s *PIXService) GetQRCodeData(input types.GetQRCodeInput) (*types.QRCode, *Response, error) {
	const path = "/api/v1/pix/outbound_pix_payments/brcodes"

	req, err := s.client.NewAPIRequest(http.MethodGet, path, input)
	if err != nil {
		return nil, nil, err
	}

	var qrcode types.QRCode
	resp, err := s.client.Do(req, &qrcode)
	if err != nil {
		return nil, resp, err
	}

	return &qrcode, resp, err
}

//ListKeys list the PIX keys of an account
func (s *PIXService) ListKeys(accountID string) ([]types.PIXKey, *Response, error){
	path := fmt.Sprintf("/api/v1/pix/%s/entries", accountID)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor   `json:"cursor"`
		Data   []types.PIXKey `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}


//ListQRCodeDynamic list the dynamic qrcodes of an account
func (s *PIXService) ListDynamicQRCodes(accountID string) ([]types.QRCodeDynamic, *Response, error){
	path := fmt.Sprintf("/api/v1/pix_payment_invoices/?account_id=%s", accountID)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.AddAccountIdHeader(req, accountID)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor   `json:"cursor"`
		Data   []types.QRCodeDynamic `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}

// DynamicQRCode make a bar code payment invoice
func (s *PIXService) DynamicQRCode(input types.QRCodeDynamicInput, idempotencyKey string) (*types.QRCodeDynamic, *Response, error) {
	path := "/api/v1/pix_payment_invoices"
	if err := input.Validate(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.AddIdempotencyHeader(req, idempotencyKey)
	if err != nil {
		return nil, nil, err
	}

	var qrcode types.QRCodeDynamic
	resp, err := s.client.Do(req, &qrcode)
	if err != nil {
		return nil, resp, err
	}

	return &qrcode, resp, err
}