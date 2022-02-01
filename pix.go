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
func (s *PixService) GetOutboundPix(id string) (*types.PIXOutBoundOutput, *Response, error) {
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
func (s *PixService) GetQRCodeData(input types.GetQRCodeInput) (*types.QRCode, *Response, error) {
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

//ListQRCodeDynamic list the dynamic qrcodes of an account
func (s *PixService) ListDynamicQRCodes(accountID string) ([]types.QRCodeDynamic, *Response, error) {
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
		Cursor types.Cursor          `json:"cursor"`
		Data   []types.QRCodeDynamic `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}

// CreateDynamicQRCode make a bar code payment invoice
func (s *PixService) CreateDynamicQRCode(input types.CreateDynamicQRCodeInput, idempotencyKey string) (*types.PIXInvoiceOutput, *Response, error) {
	const path = "/api/v1/pix_payment_invoices"

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

	var pixInvoiceOutput types.PIXInvoiceOutput
	resp, err := s.client.Do(req, &pixInvoiceOutput)
	if err != nil {
		return nil, resp, err
	}

	return &pixInvoiceOutput, resp, err
}

// CreatePedingPayment is a service used to create a pending payment.
func (s *PixService) CreatePedingPayment(input types.CreatePedingPaymentInput, idempotencyKey string) (*types.PendingPaymentOutput, *Response, error) {
	const path = "/api/v1/pix/outbound_pix_payments"

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, nil, err
	}

	err = s.client.AddIdempotencyHeader(req, idempotencyKey)
	if err != nil {
		return nil, nil, err
	}

	var pendingPaymentOutput types.PendingPaymentOutput
	resp, err := s.client.Do(req, &pendingPaymentOutput)
	if err != nil {
		return nil, resp, err
	}

	return &pendingPaymentOutput, resp, err
}

// ConfirmPedingPayment is a service used to confirm a pending payment.
func (s *PixService) ConfirmPedingPayment(input types.ConfirmPendingPaymentInput, idempotencyKey, pixID string) (*Response, error) {
	path := fmt.Sprintf("/api/v1/pix/outbound_pix_payments/%s/actions/confirm", pixID)

	req, err := s.client.NewAPIRequest(http.MethodPost, path, input)
	if err != nil {
		return nil, err
	}

	err = s.client.AddIdempotencyHeader(req, idempotencyKey)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
