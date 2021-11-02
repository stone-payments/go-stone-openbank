package openbank

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/stone-co/go-stone-openbank/types"
)

type PaymentLinkService struct {
	client *Client
}

func (s *PaymentLinkService) Get(accountID, orderID string) (types.PaymentLink, *Response, error) {
	accountID = strings.TrimSpace(accountID)
	orderID = strings.TrimSpace(orderID)

	if accountID == "" {
		return types.PaymentLink{}, nil, errors.New("account_id can't be empty")
	}

	if orderID == "" {
		return types.PaymentLink{}, nil, errors.New("order_id can't be empty")
	}

	path := fmt.Sprintf("/api/v1/payment_links/%s/orders/%s", accountID, orderID)

	req, err := s.client.NewAPIRequest(http.MethodGet, path, emptyIdempotencyKey, nil)
	if err != nil {
		return types.PaymentLink{}, nil, err
	}

	var paymentLink types.PaymentLink

	resp, err := s.client.Do(req, &paymentLink)
	if err != nil {
		return types.PaymentLink{}, resp, err
	}

	return paymentLink, resp, nil
}

func (s *PaymentLinkService) Create(input types.PaymentLinkInput) (types.PaymentLink, *Response, error) {
	path := "/api/v1/payment_links/orders"

	req, err := s.client.NewAPIRequest(http.MethodPost, path, emptyIdempotencyKey, input)
	if err != nil {
		return types.PaymentLink{}, nil, err
	}

	var paymentLink types.PaymentLink

	resp, err := s.client.Do(req, &paymentLink)
	if err != nil {
		return types.PaymentLink{}, resp, err
	}

	return paymentLink, resp, nil
}

func (s *PaymentLinkService) Cancel(orderID string, input types.PaymentLinkCancelInput) (types.PaymentLink, *Response, error) {
	orderID = strings.TrimSpace(orderID)

	if orderID == "" {
		return types.PaymentLink{}, nil, errors.New("account_id can't be empty")
	}

	if err := input.Validate(); err != nil {
		return types.PaymentLink{}, nil, err
	}

	path := fmt.Sprintf("/api/v1/payment_links/orders/%s/closed", orderID)

	req, err := s.client.NewAPIRequest(http.MethodPatch, path, emptyIdempotencyKey, input)
	if err != nil {
		return types.PaymentLink{}, nil, err
	}

	var paymentLink types.PaymentLink

	resp, err := s.client.Do(req, &paymentLink)
	if err != nil {
		return types.PaymentLink{}, resp, err
	}

	return paymentLink, resp, nil
}
