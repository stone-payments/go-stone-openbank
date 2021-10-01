package types

type TargetOrSourceAccount struct {
	Account struct {
		AccountCode string `json:"account_code"`
		BranchCode  string `json:"branch_code"`
		AccountType string `json:"account_type"`
	} `json:"account"`
	Entity      Entity      `json:"entity"`
	Institution Institution `json:"institution"`
}

type PIXOutBoundOutput struct {
	ID                       string               `json:"id"`
	AccountID                string               `json:"account_id"`
	Amount                   int                  `json:"amount"`
	CreatedAt                string               `json:"created_at"`
	Description              string               `json:"description"`
	EndToEndID               string               `json:"end_to_end_id"`
	Fee                      int                  `json:"fee"`
	RefundedAmount           int                  `json:"refunded_amount"`
	TransactionID            string               `json:"transaction_id"`
	Status                   string               `json:"status"` //currently returning: CREATED, FAILED, MONEY_RESERVED, SETTLED, REFUNDED
	Source                   TargetOrSourceAccount `json:"source"`
	Target                   TargetOrSourceAccount `json:"target"`
	CreatedBy                string               `json:"created_by"`
	FailedAt                 string               `json:"failed_at"`
	FailureReasonCode        string               `json:"failure_reason_code"`
	FailureReasonDescription string               `json:"failure_reason_description"`
	Key                      string               `json:"key"`
	MoneyReservedAt          string               `json:"money_reserved_at"`
	RequestID                string               `json:"request_id"`
	SettledAt                string               `json:"settled_at"`
	ApprovedBy               string               `json:"approved_by"`
	ApprovedAt               string               `json:"approved_at"`
}
