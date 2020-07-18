// Package jsonapi .....
package jsonapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// BaseURL for JSON Placeholder REST API
	defaultBaseURL = "https://jsonplaceholder.typicode.com/"

	//defaultMediaType = "application/vnd.api+json"
)

// Client is responsible for communicating with JSON Placeholder API.
type Client struct {
	apiKey  string
	BaseURL string

	Client *http.Client

	Accounts AccountsService
}

// NewClient func is responsible for creating a new client
func NewClient() *Client {
	c := &Client{
		Client:  &http.Client{},
		BaseURL: defaultBaseURL,
	}
	c.Accounts.client = c
	c.Accounts = c.Accounts
	return c
}

// NewRequest creates a custom new http request by setting the method, url and body
func (c *Client) NewRequest(verb, resource string, data interface{}) (*http.Request, error) {

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
	//TODO:: change defaultBaseURL to c.BaseURL
	url := c.BaseURL + resource

	req, err := http.NewRequest(verb, url, buf)
	return req, err
}

// Do invokes a 3rd Party REST API endpoint and recieves a API response back.
// The response body is the decoded inside the value pointed by data
func (c *Client) Do(ctx context.Context, req *http.Request, data interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, fmt.Errorf("error: nil context found")
	}

	req = req.WithContext(ctx)

	resp, err := c.Client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}
	defer resp.Body.Close()

	decErr := json.NewDecoder(resp.Body).Decode(data)
	if decErr == io.EOF {
		decErr = nil
	}
	if decErr != nil {
		err = decErr
	}

	// TODO:: process the reposne before returning
	// response := newResponse(resp)
	return resp, err
}