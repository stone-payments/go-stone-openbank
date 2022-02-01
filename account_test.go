package openbank

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stone-co/go-stone-openbank/types"
)

func TestAccountGet(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v1/accounts/8cbeb3d2-750f-4b14-81a1-143ad715c273", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)

		response := `
		{
				 "account_code": "403881",
  			 "branch_code": "1",
  			 "id": "8cbeb3d2-750f-4b14-81a1-143ad715c273",
  			 "owner_document": "31455351881",
  			 "owner_id": "user:08807157-f8e1-439e-a2ec-154ecb4bee13",
  			 "owner_name": "Nome da Usuária",
  			 "created_at": "2019-07-31T19:13:56Z"
		}`

		fmt.Fprint(w, response)
	})

	acct, _, err := client.Account.Get("8cbeb3d2-750f-4b14-81a1-143ad715c273")
	if err != nil {
		t.Errorf("account.Get returned error: %v", err)
	}

	expected := &types.Account{AccountCode: "403881", BranchCode: "1", OwnerDocument: "31455351881", OwnerID: "user:08807157-f8e1-439e-a2ec-154ecb4bee13",
		ID: "8cbeb3d2-750f-4b14-81a1-143ad715c273", OwnerName: "Nome da Usuária", RestrictedFeatures: false, CreatedAt: "2019-07-31T19:13:56Z"}
	if !reflect.DeepEqual(acct, expected) {
		t.Errorf("account.Get returned %+v, expected %+v", acct, expected)
	}
}
