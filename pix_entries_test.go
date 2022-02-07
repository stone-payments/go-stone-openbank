package openbank

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stone-co/go-stone-openbank/types"
)

func TestCreateEntry(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v1/pix/8cbeb3d2-750f-4b14-81a1-143ad715c273/entries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		response := `
		{
				 "verification_id": "a123456"
		}`

		fmt.Fprint(w, response)
	})

	input := types.CreatePixEntryInput{
		AccountID: "8cbeb3d2-750f-4b14-81a1-143ad715c273",
		Key:       "c1@stone.com.br",
		KeyType:   "email",
	}
	data, _, err := client.Pix.CreateEntry(input, "idempotencyKey123")
	if err != nil {
		t.Errorf("pix.CreateEntry returned error: %v", err)
	}

	expected := CreatePixEntryOutput{VerificationID: "a123456"}
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("pix.CrateEntry returned %+v, expected %+v", data, expected)
	}
}

func TestCreateEntryWithVerification(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v1/pix/8cbeb3d2-750f-4b14-81a1-143ad715c273/entries", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)

		response := `
		{
				 "id": "abcd123"
		}`

		fmt.Fprint(w, response)
	})

	input := types.CreatePixEntryInput{
		AccountID:      "8cbeb3d2-750f-4b14-81a1-143ad715c273",
		Key:            "c1@stone.com.br",
		KeyType:        "email",
		VerificationID: "a123456",
	}
	data, _, err := client.Pix.CreateEntry(input, "idempotencyKey123")
	if err != nil {
		t.Errorf("pix.CreateEntry returned error: %v", err)
	}

	expected := CreatePixEntryOutput{ID: "abcd123"}
	if !reflect.DeepEqual(data, expected) {
		t.Errorf("pix.CrateEntry returned %+v, expected %+v", data, expected)
	}
}
