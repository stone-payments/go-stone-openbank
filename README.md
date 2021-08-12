# go-stone-openbank

A Go library to connect with [Stone Open Banking API](https://docs.openbank.stone.com.br/)

## How to install

```sh
  $ go get github.com/stone-co/go-stone-openbank
```

## Example Usage

```go
package main

import (
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
		openbank.SetConsentURL(consentURL),
		openbank.WithPEMPrivateKey(pemPrivKey),
		openbank.UseSandbox(),
	//	openbank.EnableDebug(),
	)
	if err != nil {
		log.Fatal(err)
	}

	err := client.Authenticate()
	if err != nil {
		log.Fatal(err)
	}

	consentLink, err := client.ConsentLink("")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("consent_link: %s\n", consentLink)

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
		balance, _, err := client.Account.GetBalance(accounts[i].ID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Balance: %+v", balance)
 	}
}

func readFileContent(path string) []byte {
	content, _ := ioutil.ReadFile(path)
	return content
}
```

see full [example](https://github.com/stone-co/go-stone-openbank/blob/master/example/main.go)
