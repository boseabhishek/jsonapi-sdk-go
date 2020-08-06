// Package jsonapi .....
package jsonapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// BaseURL for JSON Placeholder REST API
	baseURL = "http://localhost:8080/"

	mediaType = "application/vnd.api+json"
)

// Client is responsible for communicating with JSON Placeholder API.
type Client struct {
	apiKey  string
	BaseURL string

	Client *http.Client

	Accounts *AccountsService
}

// NewClient func is responsible for creating a new client
func NewClient() *Client {
	c := &Client{
		Client:  &http.Client{},
		BaseURL: baseURL,
	}
	c.Accounts = &AccountsService{client: c}

	return c
}

// NewRequest creates a custom new http request by setting the method, url and body
func (c *Client) NewRequest(verb, resource string, data interface{}) (*http.Request, error) {

	buf := new(bytes.Buffer)
	if data != nil {
		err := json.NewEncoder(buf).Encode(data)
		if err != nil {
			return nil, err
		}
	}

	url := c.BaseURL + resource

	req, err := http.NewRequest(verb, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)

	return req, nil
}

// Perform invokes a FAke Form3 REST API endpoint and recieves a API response back.
// The response body is the decoded inside the value pointed by data
func (c *Client) Perform(ctx context.Context, req *http.Request, data interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, fmt.Errorf("nil context found")
	}

	req = req.WithContext(ctx)

	res, err := c.Client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer res.Body.Close()

	// handle resp status codes
	if c := res.StatusCode; 400 <= c && c <= 499 {
		return nil, fmt.Errorf("client error: %+v",
			&Error{Code: res.StatusCode, Message: errMessage(res)})
	}
	if c := res.StatusCode; 500 <= c && c <= 599 {
		return nil, fmt.Errorf("server error: %+v",
			&Error{Code: res.StatusCode, Message: errMessage(res)})
	}

	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return nil, fmt.Errorf("decoding response body: %v", err)
	}

	// TODO:: process the reposne before returning
	// response := newResponse(resp)
	return res, err
}

// Error custom error message
type Error struct {
	// HTTP Response Status Code
	Code int

	// Custom Response Error message
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Status Code: %v Error Message: %v", e.Code, e.Message)
}

func errMessage(res *http.Response) string {
	b, _ := ioutil.ReadAll(res.Body)
	return string(b)
}
