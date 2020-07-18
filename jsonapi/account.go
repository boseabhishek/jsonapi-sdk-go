package jsonapi

import (
	"fmt"
	"net/http"
)

type Account struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type AccountsService struct {
	client *Client
}

func (as *AccountsService) Fetch(id string) (*Account, *http.Response, error) {
	// create the resource
	u := fmt.Sprintf("posts/%s", id)

	// create a http request
	req, err := as.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(Account)
	resp, err := as.client.Do(req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
