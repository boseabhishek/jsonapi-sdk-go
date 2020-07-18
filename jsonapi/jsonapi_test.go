// Package jsonapi .....
package jsonapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	apiURL string = "/api-url"
)

func setup() (mux *http.ServeMux, client *Client, teardown func()) {
	mux = http.NewServeMux()
	srv := httptest.NewServer(mux)
	client = NewClient()

	url := srv.URL + apiURL + "/"
	client.BaseURL = url

	return mux, client, srv.Close
}

// TODO:: amend all error messages

func TestDo_BadRequest(t *testing.T) {
	mux, client, teardown := setup()
	defer teardown()

	// program the mock
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", fmt.Sprintf("%s", "."), nil)

	res, err := client.Do(req, nil)
	if err == nil {
		t.Fatalf("expected HTTP %s error, got no error", http.StatusText(http.StatusBadRequest))
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected HTTP %s error, got %s",
			http.StatusText(http.StatusBadRequest), http.StatusText(res.StatusCode))
	}
}

func TestDo_NotFound(t *testing.T) {
	mux, client, teardown := setup()
	defer teardown()

	id := "not-found-id"
	// program the mock
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", 404)
	})

	req, _ := client.NewRequest("GET", fmt.Sprintf("%s%s", client.BaseURL, id), nil)

	res, err := client.Do(req, nil)
	if err == nil {
		t.Fatalf("expected HTTP %s error, got no error", http.StatusText(http.StatusNotFound))
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected HTTP %s error, got %s",
			http.StatusText(http.StatusNotFound), http.StatusText(res.StatusCode))
	}
}

func TestOK_OK(t *testing.T) {
	mux, client, teardown := setup()
	defer teardown()

	type test struct {
		Key string
	}

	id := "present-id"
	// program the mock
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"Key":"value"}`)
	})

	req, _ := client.NewRequest("GET", fmt.Sprintf("%s%s", client.BaseURL, id), nil)
	data := new(test)

	res, err := client.Do(req, data)
	if err != nil {
		t.Errorf("expected HTTP %d success, got error: %v", http.StatusOK, err)
	}
	want := &test{"value"}

	if !reflect.DeepEqual(data, want) {
		t.Errorf("response body: %v, want: %v", data, want)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected HTTP %s success; got: %s",
			http.StatusText(http.StatusOK), http.StatusText(res.StatusCode))
	}

}
