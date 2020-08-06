package main

import (
	"context"
	"fmt"
	"go-rest/jsonapi"
)

func main() {

	client := jsonapi.NewClient()
	//d, err := client.Accounts.Fetch(context.Background(), "79c9d274-d72f-11ea-87d0-0242ac130003")
	//acs, err := client.Accounts.List(context.Background())
	/* acr := &jsonapi.AccountCreateRequest{
		CountryCode:           "TT",
		BankID:                "BANKID1",
		BankIDCode:            "BANKIDCODE",
		Bic:                   "NWBKGB22",
		Iban:                  "IBAN123",
		AccountClassification: "personal",
	}

	_, err := client.Accounts.Create(context.Background(), acr) */
	_, err := client.Accounts.Delete(context.Background(), "e2a0091e-00c9-aef7-1854-8abaf4aa5db0", 0)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}

	/* for i, ac := range acs {
		fmt.Printf("%v. %v\n", i+1, ac)
	} */

	fmt.Printf("------%+v", err)

}
