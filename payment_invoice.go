package openbank

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/stone-co/go-stone-openbank/types"
)

// PaymentInvoiceService handlers communication with Stone Openbank API
type PaymentInvoiceService struct {
	client *Client
}

// PaymentInvoice make a bar code payment invoice
func (s *PaymentInvoiceService) PaymentInvoice(input types.PaymentInvoiceInput, idempotencyKey string) (*types.PaymentInvoice, *Response, error) {
	path := "/api/v1/barcode_payment_invoices"
	if err := input.Validate(); err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewAPIRequest(http.MethodPost, path, idempotencyKey, input)
	if err != nil {
		return nil, nil, err
	}

	var paymentInvoice types.PaymentInvoice
	resp, err := s.client.Do(req, &paymentInvoice)
	if err != nil {
		return nil, resp, err
	}

	return &paymentInvoice, resp, err
}

// List returns a list of PaymentInvoices
func (s *PaymentInvoiceService) List(accountID string) ([]types.PaymentInvoice, *Response, error) {
	path := fmt.Sprintf("/api/v1/barcode_payment_invoices/?account_id=%s", accountID)
	if strings.TrimSpace(accountID) == "" {
		return nil, nil, errors.New("account_id can't be empty")
	}

	req, err := s.client.NewAPIRequest(http.MethodGet, path, emptyIdempotencyKey, nil)
	if err != nil {
		return nil, nil, err
	}

	var dataResp struct {
		Cursor types.Cursor           `json:"cursor"`
		Data   []types.PaymentInvoice `json:"data"`
	}

	resp, err := s.client.Do(req, &dataResp)
	if err != nil {
		return nil, resp, err
	}

	return dataResp.Data, resp, err
}

// Get return a PaymentInvoice
func (s *PaymentInvoiceService) Get(paymentInvoiceID string) (types.PaymentInvoice, *Response, error) {
	path := fmt.Sprintf("/api/v1/barcode_payment_invoices/%s", paymentInvoiceID)
	var paymentInvoice types.PaymentInvoice
	if strings.TrimSpace(paymentInvoiceID) == "" {
		return paymentInvoice, nil, errors.New("payment_invoice_id can't be empty")
	}

	req, err := s.client.NewAPIRequest(http.MethodGet, path, emptyIdempotencyKey, nil)
	if err != nil {
		return paymentInvoice, nil, err
	}

	resp, err := s.client.Do(req, &paymentInvoice)
	if err != nil {
		return paymentInvoice, resp, err
	}

	return paymentInvoice, resp, err
}

func (s *PaymentInvoiceService) Cancel(paymentInvoiceID string) (*Response, error) {
	paymentInvoiceID = strings.TrimSpace(paymentInvoiceID)
	if paymentInvoiceID == "" {
		return nil, errors.New("payment_invoice_id can't be empty")
	}

	path := fmt.Sprintf("/api/v1/barcode_payment_invoices/%s/cancel", paymentInvoiceID)

	req, err := s.client.NewAPIRequest(http.MethodPost, path, emptyIdempotencyKey,nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
