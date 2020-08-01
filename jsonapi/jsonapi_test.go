// Package jsonapi .....
package jsonapi

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const (
	apiURL string = "/api-url"
)

var (
	mux *http.ServeMux

	srv *httptest.Server

	client *Client

	ctx = context.TODO()
)

func setup() {
	mux = http.NewServeMux()
	srv = httptest.NewServer(mux)
	client = NewClient()

	// TODO:: explore url to get rid of trailing /
	url := srv.URL + "/"
	client.BaseURL = url

}

func teardown() {
	srv.Close()
}

func TestDo_BadRequest(t *testing.T) {
	setup()
	defer teardown()

	// program the mock
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest(http.MethodGet, ".in", nil)

	_, err := client.Do(ctx, req, nil)
	if err == nil {
		t.Fatalf("expected: HTTP %s error, got no error.", http.StatusText(http.StatusBadRequest))
	}
}

func TestDo_OK(t *testing.T) {
	setup()
	defer teardown()

	type test struct {
		Key string
	}

	mux.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"Key":"value"}`)
	})

	req, _ := client.NewRequest("GET", "/id", nil)
	data := new(test)

	res, err := client.Do(ctx, req, data)
	if err != nil {
		t.Errorf("expected: HTTP %d success, got error: %v", http.StatusOK, err)
	}
	want := &test{"value"}

	if !reflect.DeepEqual(data, want) {
		t.Errorf("response body: %v, expected: %v", data, want)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected: HTTP %s success; got: %s",
			http.StatusText(http.StatusOK), http.StatusText(res.StatusCode))
	}

}

func TestDo_nilContext(t *testing.T) {
	setup()
	defer teardown()

	req, _ := client.NewRequest("GET", fmt.Sprintf("%s", "."), nil)
	_, err := client.Do(nil, req, nil)

	// TODO:: change from reflect to normal
	if !reflect.DeepEqual(err, fmt.Errorf("error: nil context found")) {
		t.Errorf("expected `error: nil context found`")
	}
}
