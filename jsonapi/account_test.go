package jsonapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

/* func TestList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("expected method: %v, got: %v", m, r.Method)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{
			"userId": 1,
			"id": 1,
			"title": "title 1",
			"body": "body 1"
		  }]`)
	})

	account, _, err := client.Accounts.List(ctx)
	if err != nil {
		t.Fatalf("Accounts.Fetch returned error: %v", err)
	}

	want := &Account{
		Data: Data{
			Type:           DataType,
			ID:             "some-id",
			OrganisationID: "some-id",
			Version:        0,
			Attributes: Attributes{
				Country: "TT", BankID: "BANKID",
				BankIDCode: "BANKIDCODE", Bic: "vbgb",
				AccountClassification: "Personal"}}}

	if !reflect.DeepEqual(account, want) {
		t.Errorf("Account data fetched: got=%#v\nwant=%#v", account, want)
	}

} */

func TestFetch(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/organisation/accounts/some-id", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("expected method: %v, got: %v", m, r.Method)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"data": {
				"type": "accounts",
				"id": "some-id",
				"organisation_id": "some-id",
				"version": 0,
				"attributes": {
					"country": "TT",
					"bank_id": "BANKID",
					"bank_id_code": "BANKIDCODE",
					"bic": "vbgb",
					"account_classification": "Personal"
				}
			}
		}`)
	})

	account, _, err := client.Accounts.Fetch(ctx, "some-id")
	if err != nil {
		t.Fatalf("Accounts.Fetch returned error: %v", err)
	}

	want := &Account{
		Data: Data{
			Type:           DataType,
			ID:             "some-id",
			OrganisationID: "some-id",
			Version:        0,
			Attributes: Attributes{
				Country: "TT", BankID: "BANKID",
				BankIDCode: "BANKIDCODE", Bic: "vbgb",
				AccountClassification: "Personal"}}}

	if !reflect.DeepEqual(account, want) {
		t.Errorf("Account data fetched: got=%#v\nwant=%#v", account, want)
	}

}

func TestFetch_BlankId(t *testing.T) {
	setup()
	defer teardown()

	_, _, err := client.Accounts.Fetch(ctx, "")

	// TODO: custom error handling must be implemented later
	// avoid reflect.DeepEqual
	if !reflect.DeepEqual(err, fmt.Errorf("id can't be blank or empty")) {
		t.Fatalf("Accounts.Fetch err recieved %+v", err)
	}

}

func TestFetch_NonMatchingJsonBodyRecieved(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/organisation/accounts/some-id", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("expected method: %v, got: %v", m, r.Method)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"a": "b",
			"c": "d"
		  }`)
	})

	account, _, err := client.Accounts.Fetch(ctx, "some-id")
	if err != nil {
		t.Fatalf("Accounts.Fetch error: %v", err)
	}

	want := &Account{
		Data: Data{
			Type:           DataType,
			ID:             "some-id",
			OrganisationID: "some-id",
			Version:        0,
			Attributes: Attributes{
				Country: "TT", BankID: "BANKID",
				BankIDCode: "BANKIDCODE", Bic: "vbgb",
				AccountClassification: "Personal"}}}

	if reflect.DeepEqual(account, want) {
		t.Errorf("Accounts.Fetch non-matching json recieved got=%#v\nwant=%#v", account, want)
	}
}
