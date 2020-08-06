package jsonapi

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/organisation/accounts", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("expected method: %v, got: %v", m, r.Method)
		}
		w.WriteHeader(http.StatusCreated)
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

	acr := &AccountCreateRequest{ 
		CountryCode:           "TT",
		BankID:                "BANKID",
		BankIDCode:            "BANKIDCODE",
		Bic:                   "vbgb",
		Iban:                  "IBAN123",
		AccountClassification: "personal",
	}

	account, err := client.Accounts.Create(ctx, acr)
	if err != nil {
		t.Fatalf("Accounts.Create returned error: %v", err)
	}

	want := &Account{
		Data: Data{
			Type:           AccountDataType,
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

func TestCreate_NoCreatedStatusReturned(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/organisation/accounts", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("expected method: %v, got: %v", m, r.Method)
		}
		w.WriteHeader(http.StatusAccepted)
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

	acr := &AccountCreateRequest{
		CountryCode:           "TT",
		BankID:                "BANKID",
		BankIDCode:            "BANKIDCODE",
		Bic:                   "vbgb",
		Iban:                  "IBAN123",
		AccountClassification: "personal",
	}

	_, err := client.Accounts.Create(ctx, acr)

	if !errors.Is(err, &Error{Type: HttpErrorType, Code: 202, Message: "Status Code for Create response must be 201"}) {
		t.Fatalf("Accounts.Fetch err recieved %+v", err)
	}

}

func TestList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/organisation/accounts", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("expected method: %v, got: %v", m, r.Method)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"data": [{
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
			}]
		}`)
	})

	accounts, err := client.Accounts.List(ctx)
	if err != nil {
		t.Fatalf("Accounts.List returned error: %v", err)
	}

	want := &AccountList{
		Data: []Data{{
			Type:           AccountDataType,
			ID:             "some-id",
			OrganisationID: "some-id",
			Version:        0,
			Attributes: Attributes{
				Country: "TT", BankID: "BANKID",
				BankIDCode: "BANKIDCODE", Bic: "vbgb",
				AccountClassification: "Personal"}}}}

	if !reflect.DeepEqual(accounts, want) {
		t.Errorf("Account data fetched: got=%#v\nwant=%#v", accounts, want)
	}

}

func TestList_NonMatchingJsonBodyRecieved(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/organisation/accounts", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("expected method: %v, got: %v", m, r.Method)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"a": "b",
			"c": "d"
		  }`)
	})

	accounts, err := client.Accounts.List(ctx)
	if err != nil {
		t.Fatalf("Accounts.List returned error: %v", err)
	}

	want := &AccountList{
		Data: []Data{{
			Type:           AccountDataType,
			ID:             "some-id",
			OrganisationID: "some-id",
			Version:        0,
			Attributes: Attributes{
				Country: "TT", BankID: "BANKID",
				BankIDCode: "BANKIDCODE", Bic: "vbgb",
				AccountClassification: "Personal"}}}}

	if reflect.DeepEqual(accounts, want) {
		t.Errorf("Accounts.List json recieved=%#v\n must not match with expected=%#v", accounts, want)
	}
}

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

	account, err := client.Accounts.Fetch(ctx, "some-id")
	if err != nil {
		t.Fatalf("Accounts.Fetch returned error: %v", err)
	}

	want := &Account{
		Data: Data{
			Type:           AccountDataType,
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

	_, err := client.Accounts.Fetch(ctx, "")

	if !errors.Is(err, &Error{Type: UserErrorType, Message: "id can't be blank or empty"}) {
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

	account, err := client.Accounts.Fetch(ctx, "some-id")
	if err != nil {
		t.Fatalf("Accounts.Fetch error: %v", err)
	}

	want := &Account{
		Data: Data{
			Type:           AccountDataType,
			ID:             "some-id",
			OrganisationID: "some-id",
			Version:        0,
			Attributes: Attributes{
				Country: "TT", BankID: "BANKID",
				BankIDCode: "BANKIDCODE", Bic: "vbgb",
				AccountClassification: "Personal"}}}

	if reflect.DeepEqual(account, want) {
		t.Errorf("Accounts.Fetch json recieved=%#v\n must not match with expected=%#v", account, want)
	}
}
