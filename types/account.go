package types

//Account represents a Stone PaymentAccount
type Account struct {
	AccountCode        string `json:"account_code"`
	BranchCode         string `json:"branch_code"`
	ID                 string `json:"id"`
	OwnerDocument      string `json:"owner_document"`
	OwnerID            string `json:"owner_id"`
	OwnerName          string `json:"owner_name"`
	RestrictedFeatures bool   `json:"restricted_features"`
	CreatedAt          string `json:"created_at,omitempty"`
}

//Balance represents a Stone PaymentAccount Balance
type Balance struct {
	Balance          int `json:"balance"`
	BlockedBalance   int `json:"blocked_balance"`
	ScheduledBalance int `json:"scheduled_balance"`
}

type Statement struct {
	ID                      string `json:"id"`
	Type                    string `json:"type"`
	Amount                  int    `json:"amount"`
	BalanceAfter            int    `json:"balance_after,omitempty"`
	BalanceBefore           int    `json:"balance_before,omitempty"`
	CreatedAt               string `json:"created_at,omitempty"`
	UpdatedAt               string `json:"updated_at,omitempty"`
	Status                  string `json:"status,omitempty"`
	Operation               string `json:"operation,omitempty"`
	OperationID             string `json:"operation_id,omitempty"`
	Description             string `json:"description,omitempty"`
	OperationAmount         int    `json:"operation_amount,omitempty"`
	FeeAmount               int    `json:"fee_amount,omitempty"`
	RefundReasonCode        string `json:"refund_reason_code,omitempty"`
	RefundReasonDescription string `json:"refund_reason_description,omitempty"`
	OriginalOperationID     string `json:"original_operation_id,omitempty"`
	RefundedAt              string `json:"refunded_at,omitempty"`
	Barcode                 string `json:"barcode,omitempty"`

	CardNetworkCode string `json:"card_network_code,omitempty"`
	CardNetworkName string `json:"card_network_name,omitempty"`
	CardType        string `json:"card_type,omitempty"`
	IsPrepayment    bool   `json:"is_prepayment,omitempty"`

	Details struct {
		BankName         string `json:"bank_name,omitempty"`
		RecipientCpfCnpj string `json:"recipient_cpf_cnpj,omitempty"`
		RecipientName    string `json:"recipient_name,omitempty"`
		WritableLine     string `json:"writable_line,omitempty"`
		ExpirationDate   string `json:"expiration_date,omitempty"`
	} `json:"details,omitempty"`

	CounterParty CounterParty `json:"counter_party,omitempty"`

	DelayedToNextBusinessDay bool `json:"delayed_to_next_business_day,omitempty"`
}

type CounterParty struct {
	Account struct {
		Institution     string `json:"institution,omitempty"`
		InstitutionName string `json:"institution_name,omitempty"`
		AccountCode     string `json:"account_code,omitempty"`
		BranchCode      string `json:"branch_code,omitempty"`
		AccountType     string `json:"account_type,omitempty"`
	} `json:"account"`
	Entity Entity `json:"entity"`
}

type Fee struct {
	Amount                      int    `json:"amount"`
	FeeType                     string `json:"fee_type"`
	BillingExemptionParticipant bool   `json:"billing_exemption_participant"`
	OriginalFee                 int    `json:"original_fee"`
	MaxFreeTransfers            int    `json:"max_free_transfers"`
	RemainingFreeTransfers      int    `json:"remaining_free_transfers"`
}

func ListFeeTypes() []string {
	return []string{
		"internal_transfer",
		"external_transfer",
		"barcode_payment",
		"outbond_stone_prepaid_card_wirhdrawal",
		"barcode_payment_invoice",
	}
}
