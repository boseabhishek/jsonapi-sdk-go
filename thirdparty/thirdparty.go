// Package thirdparty .....
package thirdparty

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	// BaseURL for REST API
	defaultBaseURL = "https://jsonplaceholder.typicode.com/"

	//defaultMediaType = "application/vnd.api+json"
)

// Client is responsible for communicating with JSON Placeholder API.
type Client struct {
	Client  *http.Client
	BaseURL *url.URL
}

// NewClient func is responsible for creating a new thirdparty client
func NewClient() *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		Client:  &http.Client{},
		BaseURL: baseURL,
	}
}

// NewRequest creates a custom new http request by setting the method, url and body
func (fc *Client) NewRequest(verb, url string, data interface{}) (*http.Request, error) {
	//TODO:: try url parse (see also notes)
	// fix bytes.NewBuffer(jsonReq)
	var buf io.ReadWriter
	if data != nil { // mostly for POST, PUT etc.
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(data)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(verb, defaultBaseURL+url, buf)
	return req, err
}

// Do invokes a 3rd Party REST API endpoint and recieves a API response back.
// The response body is the decoded inside the value pointed by data
// TODO:: change *http.Response to custom Response
func (fc *Client) Do(req *http.Request, data interface{}) (*http.Response, error) {

	resp, err := fc.Client.Do(req)
	if err != nil {
		// TODO:: process err
		// maybe introduce ctx context.Context as param and then do as below:
		// select {
		// case <-ctx.Done():
		// 	return nil, ctx.Err()
		// default:
		return nil, err
	}
	defer resp.Body.Close()

	decErr := json.NewDecoder(resp.Body).Decode(data)
	if decErr == io.EOF {
		decErr = nil // ignore EOF errors caused by empty response body
	}
	if decErr != nil {
		err = decErr
	}

	// TODO:: process the reposne before returning
	// response := newResponse(resp)
	return resp, err
}
