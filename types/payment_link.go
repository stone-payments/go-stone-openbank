package types

import "encoding/json"

type PaymentLink struct {
	ID        string                `json:"id"`
	Amount    int                   `json:"amount"`
	Checkouts []PaymentLinkCheckout `json:"checkouts"`
	Closed    bool                  `json:"closed"`
	Code      string                `json:"code"`
	Currency  string                `json:"currency"`
	Customer  PaymentLinkCustomer   `json:"customer"`
	Items     []PaymentLinkItem     `json:"items"`
	SessionID string                `json:"session_id"`
	Status    string                `json:"status"`
	CreatedAt string                `json:"created_at"`
	UpdatedAt string                `json:"updated_at"`
}

type PaymentLinkCheckout struct {
	ID                          string                `json:"id"`
	AcceptedMultiPaymentMethods []string              `json:"accepted_multi_payment_methods"`
	AcceptedPaymentMethods      []string              `json:"accepted_payment_methods"`
	Amount                      int                   `json:"amount"`
	BillingAddress              json.RawMessage       `json:"billing_address"`
	BillingAddressEditable      bool                  `json:"billing_address_editable"`
	CreditCard                  PaymentLinkCreditCard `json:"credit_card"`
	Currency                    string                `json:"currency"`
	Customer                    PaymentLinkCustomer   `json:"customer"`
	CustomerEditable            bool                  `json:"customer_editable"`
	ExpiresAt                   string                `json:"expires_at"`
	Metadata                    json.RawMessage       `json:"metadata"`
	PaymentURL                  string                `json:"payment_url"`
	RequiredFields              []string              `json:"required_fields"`
	Shippable                   bool                  `json:"shippable"`
	SkipCheckoutSuccessPage     bool                  `json:"skip_checkout_success_page"`
	Status                      string                `json:"status"`
	SuccessURL                  string                `json:"success_url"`
	CreatedAt                   string                `json:"created_at"`
	UpdatedAt                   string                `json:"updated_at"`
}

type PaymentLinkCreditCard struct {
	Authentication PaymentLinkCreditCardAuth          `json:"authentication"`
	Capture        bool                               `json:"capture"`
	Installments   []PaymentLinkCreditCardInstallment `json:"installments"`
}

type PaymentLinkCreditCardAuth struct {
	ThreedSecure json.RawMessage `json:"threed_secure"`
	Type         string          `json:"type"`
}

type PaymentLinkCreditCardInstallment struct {
	Number int `json:"number"`
	Total  int `json:"total"`
}

type PaymentLinkCustomer struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Delinquent bool            `json:"delinquent"`
	Phones     json.RawMessage `json:"phones"`
	CreatedAt  string          `json:"created_at"`
	UpdatedAt  string          `json:"updated_at"`
}

type PaymentLinkItem struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
