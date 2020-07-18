package jsonapi

import (
	"context"
	"fmt"
	"net/http"
)

// Account represents an account from JsonAPI

// An Account represents a bank account that is registered with Form3
type Account struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// AccountsService conducts the conversation with account related
// resources of JsonAPI
type AccountsService struct {
	client *Client
}

// Fetch retrives the account information
func (as *AccountsService) Fetch(ctx context.Context, id string) (*Account, *http.Response, error) {
	// create the resource
	u := fmt.Sprintf("posts/%s", id)

	// create a http request
	req, err := as.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(Account)
	resp, err := as.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
