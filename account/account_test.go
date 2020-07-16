// Package account TODO:: account must be renamed to form3
package account

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	fakeBaseURL string = "/some-version/some-url"
)

func setup() (mux *http.ServeMux, client *Client, teardown func()) {
	mux = http.NewServeMux()
	srv := httptest.NewServer(mux)
	client = NewClient()

	url, _ := url.Parse(srv.URL)
	client.BaseURL = url

	return mux, client, srv.Close
}

func TestDo_NotOK(t *testing.T) {
	mux, client, teardown := setup()
	defer teardown()

	mux.HandleFunc("/b", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", fmt.Sprintf("%s/", client.BaseURL), nil)

	resp, err := client.Do(req, nil)
	if err == nil {
		t.Fatal("Expected HTTP 400 error, got no error.")
	}

	if resp.StatusCode != 400 {
		t.Errorf("Expected HTTP 400 error, got %d status code.", resp.StatusCode)
	}
}

/*
func TestOK_OK(t *testing.T) {
	json := `{
		"key": "value"
	}`
	srv := NewMockServer(http.StatusOK, json)
	defer srv.Close()

	fakeClient := &account.Client{
		Client: srv.Client(),
		URL:    srv.URL + "/get-list/" + "some-id-1",
	}

	res, err := fakeClient.Fetch()
	if err != nil {
		log.Fatal(err)
	}

	if json != string(res) {
		t.Errorf("response body doesn't match: want %s but got %s", json, string(res))
	}

}
*/
