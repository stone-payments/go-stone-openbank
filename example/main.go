package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/uuid"

	openbank "github.com/stone-co/go-stone-openbank"
	"github.com/stone-co/go-stone-openbank/types"
)

func main() {
	clientID := os.Getenv("STONE_CLIENT_ID")
	privKeyPath := os.Getenv("STONE_PRIVATE_KEY")
	consentURL := os.Getenv("STONE_CONSENT_REDIRECT_URL")

	pemPrivKey := readFileContent(privKeyPath)

	client, err := openbank.NewClient(
		openbank.WithClientID(clientID),
		openbank.WithPEMPrivateKey(pemPrivKey),
		openbank.SetConsentURL(consentURL),
		openbank.UseSandbox(),
		//	openbank.EnableDebug(),
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Authenticate(); err != nil {
		log.Fatal(err)
	}

	consentLink, err := client.ConsentLink("")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nconsent_link: %s\n", consentLink)

	// returns institutions
	allinstitutions, _, err := client.Institution.List(openbank.AllInstitutions)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(allinstitutions), allinstitutions[0])

	// returns institutions participating in the SPI. Useful for PIX operations
	SPIinstitutions, _, err := client.Institution.List(openbank.SPIParticipants)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(SPIinstitutions), SPIinstitutions[0])

	// returns institutions participating in the STR. Useful for TED operations
	STRinstitutions, _, err := client.Institution.List(openbank.STRParticipants)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(len(STRinstitutions), STRinstitutions[0])

	// return institution by code or ISPB code
	institution, _, err := client.Institution.Get(SPIinstitutions[0].ISPBCode)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(institution)

	accounts, _, err := client.Account.List()
	if err != nil {
		log.Fatal(err)
	}
	for i := range accounts {
		log.Printf("acc[%d]: %v\n\n", i, accounts[i])
		acc, _, err := client.Account.Get(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Detailed account: %+v", acc)

		balance, _, err := client.Account.GetBalance(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Balance: %+v", balance)

		statement, _, err := client.Account.GetStatement(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Statement: %+v", statement)

		fees, _, err := client.Account.ListFees(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("AllFees: %+v", fees)

		fee, _, err := client.Account.GetFees(accounts[i].ID, "internal_transfer")
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Fee: %+v", fee)

		//Internal DryRun Transfer
		transfInput := types.TransferInput{
			AccountID: accounts[i].ID,
			Amount:    100,
			Target: types.Target{
				Account: types.TransferAccount{
					AccountCode: "334201",
				},
			},
		}

		idempotencyKey := uuid.New().String()
		transfer, _, err := client.Transfer.DryRunTransfer(transfInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Transfer(dry-run): %+v", transfer)

		//Internal Transfer
		transfer, _, err = client.Transfer.Transfer(transfInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Transfer: %+v", transfer)

		//External DryRun Transfer
		transfExtInput := types.TransferInput{
			AccountID: accounts[i].ID,
			Amount:    100,
			Target: types.Target{
				Account: types.TransferAccount{
					AccountCode:     "1234",
					BranchCode:      "7032",
					InstitutionCode: "001",
				},
				Entity: types.Entity{
					Name:         "James Bond",
					Document:     "00700700700",
					DocumentType: "cpf",
				},
			},
		}
		transfer, _, err = client.Transfer.DryRunTransfer(transfExtInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("External Transfer(dry-run): %+v", transfer)

		//External  Transfer
		transfer, _, err = client.Transfer.Transfer(transfExtInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("External Transfer: %+v", transfer)

		//List Internal Transfers
		internalTransfers, _, err := client.Transfer.ListInternal(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Internal Transfers: %+v\n", internalTransfers)
		for i, t := range internalTransfers {
			//Get an internal transfer
			transfer, _, err := client.Transfer.GetInternal(t.ID)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Internal Transfer[%d]: %+v\n", i, transfer)
		}

		//List External Transfers
		externalTransfers, _, err := client.Transfer.ListExternal(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("External Transfers: %+v\n", externalTransfers)
		for i, t := range externalTransfers {
			//Get an external transfer
			transfer, _, err := client.Transfer.GetExternal(t.ID)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("External Transfer[%d]: %+v\n", i, transfer)
		}

		//Schedule and Cancel an internal Transfer
		ScheduleAndCancelTransfer(accounts[i].ID, client)

		//Payment Invoice
		paymentInvoiceInput := types.PaymentInvoiceInput{
			AccountID:      accounts[i].ID,
			Amount:         5000,
			ExpirationDate: time.Now().Format("2006-01-02"),
			InvoiceType:    "deposit",
		}

		// Make Payment Invoice
		paymentInvoice, _, err := client.PaymentInvoice.PaymentInvoice(paymentInvoiceInput, idempotencyKey)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Payment Invoice: %+v", paymentInvoice)

		//List Payment Invoices
		paymentInvoices, _, err := client.PaymentInvoice.List(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Payment Invoices: %+v\n", paymentInvoices)
		for i, t := range paymentInvoices {
			//Get a payment invoice
			invoice, _, err := client.PaymentInvoice.Get(t.ID)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Payment Invoice[%d]: %+v\n", i, invoice)
		}
	}
}

func readFileContent(path string) []byte {
	content, _ := ioutil.ReadFile(path)
	return content
}

func ScheduleAndCancelTransfer(accID string, client *openbank.Client) {
	transfInput := types.TransferInput{
		AccountID:   accID,
		Amount:      100,
		ScheduledTo: "2020-03-25",
		Target: types.Target{
			Account: types.TransferAccount{
				AccountCode: "334201",
			},
		},
	}

	idempotencyKey := uuid.New().String()
	transfer, _, err := client.Transfer.Transfer(transfInput, idempotencyKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transfer: %+v", transfer)

	//Check transfer status
	intTransfer, _, err := client.Transfer.GetInternal(transfer.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Scheduled Transfer Status: %s\n", intTransfer.Status)

	//Cancel transfer
	resp, err := client.Transfer.CancelInternal(transfer.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response from Cancel Transfer: %+v\n", resp.Response)

	//Check if transfer was canceled
	canceledTransfer, resp, err := client.Transfer.GetInternal(transfer.ID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Transfer Status: %s\n", canceledTransfer.Status)
}
