package jsonapi

import (
	"context"
	"fmt"
	"net/http"
)

// Account represents an account from JsonAPI
//
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

// Fetch retrives the account information given an id
// TODO: handle response here and get rid of response
func (as *AccountsService) Fetch(ctx context.Context, id string) (*Account, *http.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id can't be blank or empty")
	}

	// create the resource
	u := fmt.Sprintf("posts/%s", id)

	// create a http request
	req, err := as.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	ac := new(Account)
	resp, err := as.client.Perform(ctx, req, ac)
	if err != nil {
		return nil, resp, err
	}

	return ac, resp, nil
}

// List retrives all the accounts
// TODO: add ListOptions kind of thinhg with page and limit etc
// TODO: handle response here and get rid of response
func (as *AccountsService) List(ctx context.Context) ([]*Account, *http.Response, error) {

	// create the resource
	u := fmt.Sprintf("posts")

	// create a http request
	req, err := as.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var acs []*Account

	resp, err := as.client.Perform(ctx, req, &acs)
	if err != nil {
		return nil, resp, err
	}

	return acs, resp, nil
}
