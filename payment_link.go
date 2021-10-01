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

	req, err := s.client.NewAPIRequest(http.MethodGet, path, nil)
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
