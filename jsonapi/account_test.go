package jsonapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFetch(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/posts/1", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodGet; m != r.Method {
			t.Errorf("expected method: %v, got: %v", m, r.Method)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			"userId": 1,
			"id": 1,
			"title": "title 1",
			"body": "body 1"
		  }`)
	})

	account, _, err := client.Accounts.Fetch(ctx, "1")
	if err != nil {
		t.Fatalf("Accounts.Fetch returned error: %v", err)
	}

	want := &Account{UserID: 1,
		ID: 1, Title: "title 1", Body: "body 1"}

	if !reflect.DeepEqual(account, want) {
		t.Errorf("Account data fetched: got=%#v\nwant=%#v", account, want)
	}

}

func TestFetch_BlankId(t *testing.T) {
	setup()
	defer teardown()

	_, _, err := client.Accounts.Fetch(ctx, "")

	expectedErr := fmt.Errorf("id can't be blank or empty")

	if err.Error() != expectedErr.Error() {
		t.Fatalf("Accounts.Fetch expected error: %v got: %v", expectedErr, err)
	}

}
