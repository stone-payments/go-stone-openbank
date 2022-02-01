package types

import (
	"errors"
	"time"
)

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
	ID                       string                `json:"id"`
	AccountID                string                `json:"account_id"`
	Amount                   int                   `json:"amount"`
	CreatedAt                string                `json:"created_at"`
	Description              string                `json:"description"`
	EndToEndID               string                `json:"end_to_end_id"`
	Fee                      int                   `json:"fee"`
	RefundedAmount           int                   `json:"refunded_amount"`
	TransactionID            string                `json:"transaction_id"`
	Status                   string                `json:"status"` //currently returning: CREATED, FAILED, MONEY_RESERVED, SETTLED, REFUNDED
	Source                   TargetOrSourceAccount `json:"source"`
	Target                   TargetOrSourceAccount `json:"target"`
	CreatedBy                string                `json:"created_by"`
	FailedAt                 string                `json:"failed_at"`
	FailureReasonCode        string                `json:"failure_reason_code"`
	FailureReasonDescription string                `json:"failure_reason_description"`
	Key                      string                `json:"key"`
	MoneyReservedAt          string                `json:"money_reserved_at"`
	RequestID                string                `json:"request_id"`
	SettledAt                string                `json:"settled_at"`
	ApprovedBy               string                `json:"approved_by"`
	ApprovedAt               string                `json:"approved_at"`
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
	Revision          int                    `json:"revision,omitempty"`
	RequestedForPayer string                 `json:"request_for_payer,omitempty"`
	Status            string                 `json:"status,omitempty"`
	TxnID             string                 `json:"transaction_id,omitempty"`
	Amount            int                    `json:"amount,omitempty"`
	AdditionalData    []QRCodeAdditionalData `json:"additional_data,omitempty"`
}

type PIXInvoiceOutput struct {
	ID              string                 `json:"id"`
	AccountID       string                 `json:"account_id"`
	Status          string                 `json:"status"`
	Key             string                 `json:"key"`
	KeyType         string                 `json:"key_type"`
	TransactionID   string                 `json:"transaction_id"`
	Amount          int                    `json:"amount"`
	AdditionalData  []QRCodeAdditionalData `json:"additional_data"`
	RequestID       string                 `json:"request_id"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
	PaidAt          time.Time              `json:"paid_at,omitempty"`
	CancelleddAt    time.Time              `json:"cancelled_at,omitempty"`
	CreatedBy       string                 `json:"created_by"`
	LastUpdatedBy   string                 `json:"last_updated_by"`
	Expiration      int                    `json:"expiration"`
	QrCodeContent   string                 `json:"qr_code_content"`
	QrCodeImage     string                 `json:"qr_code_image"`
	RequestForPayer string                 `json:"request_for_payer"`
}

type CreatePedingPaymentInput struct {
	AccountID     string                 `json:"account_id,omitempty"`
	Amount        int                    `json:"amount,omitempty"`
	Description   string                 `json:"description,omitempty"`
	TransactionID string                 `json:"transaction_id,omitempty"`
	Key           string                 `json:"key,omitempty"`
	Source        *TargetOrSourceAccount `json:"source,omitempty"`
}
type PendingPaymentOutput struct {
	ID                       string      `json:"id"`
	AccountID                string      `json:"account_id"`
	Amount                   int         `json:"amount"`
	CreatedAt                time.Time   `json:"created_at"`
	CreatedBy                string      `json:"created_by"`
	Description              string      `json:"description"`
	TransactionID            string      `json:"transaction_id"`
	Key                      string      `json:"key"`
	EndToEndID               string      `json:"end_to_end_id"`
	FailedAt                 interface{} `json:"failed_at"`
	FailureReasonCode        interface{} `json:"failure_reason_code"`
	FailureReasonDescription interface{} `json:"failure_reason_description"`
	MoneyReservedAt          interface{} `json:"money_reserved_at"`
	RefundedAmount           int         `json:"refunded_amount"`
	RequestID                string      `json:"request_id"`
	SettledAt                interface{} `json:"settled_at"`
	Source                   struct {
		Account struct {
			AccountCode string `json:"account_code"`
			AccountType string `json:"account_type"`
			BranchCode  string `json:"branch_code"`
		} `json:"account"`
		Entity struct {
			Document     string `json:"document"`
			DocumentType string `json:"document_type"`
			Name         string `json:"name"`
		} `json:"entity"`
		Institution struct {
			Ispb string `json:"ispb"`
			Name string `json:"name"`
		} `json:"institution"`
	} `json:"source"`
	Status string `json:"status"`
	Target struct {
		Account struct {
			AccountCode string `json:"account_code"`
			AccountType string `json:"account_type"`
			BranchCode  string `json:"branch_code"`
		} `json:"account"`
		Entity struct {
			Document     string `json:"document"`
			DocumentType string `json:"document_type"`
			Name         string `json:"name"`
		} `json:"entity"`
		Institution struct {
			Ispb string `json:"ispb"`
			Name string `json:"name"`
		} `json:"institution"`
	} `json:"target"`
}

type ConfirmPendingPaymentInput struct {
	Amount              int    `json:"amount"`
	Description         string `json:"description"`
	AddTargetToContacts bool   `json:"add_target_to_contacts"`
}

type QRCodeAdditionalData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CreateDynamicQRCodeInput struct {
	Amount          int                    `json:"amount,omitempty"`
	AccountID       string                 `json:"account_id"`
	Key             string                 `json:"key,omitempty"`
	TransactionID   string                 `json:"transaction_id,omitempty"`
	AdditionalData  []QRCodeAdditionalData `json:"additional_data,omitempty"`
	RequestForPayer string                 `json:"request_for_payer"`
}

type Customer struct {
	Name     string `json:"name"`
	Document string `json:"document"`
}

type QRCodeStatic struct {
	Key    string `json:"key,omitempty"`
	Type   string `json:"phone,omitempty"`
	TxnID  string `json:"transaction_id,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

type BeneficiaryAccount struct {
	BranchCode  string `json:"branch_code"`
	AccountCode string `json:"account_code"`
	AccountType string `json:"account_type"`
	CreatedAt   string `json:"created_at"`
}

type BeneficiaryEntity struct {
	Name         string `json:"name"`
	DocumentType string `json:"document_type"`
	Document     string `json:"document"`
}

type PixEntry struct {
	ID                 string              `json:"id"`
	Key                string              `json:"key"`
	KeyType            string              `json:"key_type"`
	Status             string              `json:"status"`
	AccountID          string              `json:"account_id"`
	ParticipantISPB    string              `json:"participant_ispb"`
	BeneficiaryAccount *BeneficiaryAccount `json:"beneficiary_account"`
	BeneficiaryEntity  *BeneficiaryEntity  `json:"beneficiary_entity"`
}

type CreateEntryInput struct {
	Key              string `json:"key"`
	KeyType          string `json:"key_type"`
	AccountID        string `json:"account_id"`
	ParticipantISPB  string `json:"participant_ispb"`
	VerificationID   string `json:"verification_id,omitempty"`
	VerificationCode string `json:"verification_code,omitempty"`
}

func (p CreateDynamicQRCodeInput) Validate() error {
	if p.AccountID == "" {
		return errors.New("account_id can't be empty")
	}

	if p.Key == "" {
		return errors.New("key can't be empty")
	}

	if p.TransactionID == "" {
		return errors.New("transaction_id can't be empty")
	}
	if len(p.TransactionID) < 26 || len(p.TransactionID) > 35 {
		return errors.New("transaction_id size need to be between 26 and 35 caracteres")
	}

	return nil
}
