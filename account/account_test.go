// TODO:: account must be renamed to form3
package account

import (
	ac "go-rest/account"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	fakeBaseURL string = "/some-version/some-url"
)

func setup() (client *ac.Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	// TODO:: change to a handler impl (see notes Amit Saha)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Success!"))
	}))

	// client is a fake client used for tested and must be used for test server `server`
	client = ac.NewClient()

	url, _ := url.Parse(server.URL + fakeBaseURL + "/")
	// overriding the baseURL with a fake one
	client.BaseURL = url

	return client, mux, server.URL, server.Close
}

func TestDo_NotOK(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := client.NewRequest("GET", ".", nil)

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
