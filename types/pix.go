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

type GetQRCodeInput struct {
	BRCode       string `json:"brcode"`
	OwnerAccount string `json:"owner_account,omitempty"`
	Date         string `json:"payment_date,omitempty"`
}

type QRCode struct {
	Type    string        `json:"type"`
	Static  QRCodeStatic  `json:"static,omitempty"`
	Dynamic QRCodeDynamic `json:"dynamic,omitempty"`
}

type QRCodeDynamic struct {
	CreatedAt   string `json:"created_at,omitempty"`
	RequestedAt string `json:"requested_at,omitempty"`
	Expiration  int    `json:"expiration,omitempty"`
	Key         string `json:"key,omitempty"`
	Customer    struct {
		Name         string `json:"name"`
		Document     string `json:"document"`
		DocumentType string `json:"document_type"`
	} `json:"customer"`
	Revision          int    `json:"revision,omitempty"`
	RequestedForPayer string `json:"request_for_payer,omitempty"`
	Status            string `json:"status,omitempty"`
	TxnID             string `json:"transaction_id,omitempty"`
	Amount            int    `json:"amount,omitempty"`
}

type QRCodeStatic struct {
	Key    string `json:"key,omitempty"`
	Type   string `json:"phone,omitempty"`
	TxnID  string `json:"transaction_id,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

type BeneficiaryAccount struct {
	BranchCode string `json:"branch_code"`
	AccountCode string `json:"account_code"`
	AccountType string `json:"account_type"`
	CreatedAt string `json:"created_at"`
}

type BeneficiaryEntity struct {
	Name string `json:"name"`
	DocumentType string `json:"document_type"`
	Document string `json:"document"`
}

type KeyPIX struct {
	ID string `json:"id"`
	Key string `json:"key"`
	KeyType string `json:"key_type"`
	Status string `json:"status"`
	AccountID string `json:"account_id"`
	ParticipantISPB string `json:"participant_ispb"`
	BeneficiaryAccount *BeneficiaryAccount `json:"beneficiary_account"`
	BeneficiaryEntity *BeneficiaryEntity `json:"beneficiary_entity"`
}
