// Package account TODO:: account must be renamed to form3
package account

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	fakeBaseURL string = "/some-version/some-url"
)

func mockHandler() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	return
}
func TestDo_NotOK(t *testing.T) {

	srv := httptest.NewServer(mockHandler())
	defer srv.Close()

	// client is a fake client used for tested and must be used for test server `server`
	client := NewClient()

	//url, _ := url.Parse(srv.URL + fakeBaseURL + "/")
	// overriding the baseURL with a fake one
	//client.BaseURL = url

	req, _ := client.NewRequest("GET", fmt.Sprintf("%s/", srv.URL), nil)

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
