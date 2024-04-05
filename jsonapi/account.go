package jsonapi

import (
	"context"
	"fmt"
	"net/http"
)

// Todo represents the json shape from JsonAPI
type Todo struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// TodosService conducts the conversation with Todo related
// resources of JsonAPI
type TodosService struct {
	client *Client
}

// Fetch retrives the Todo information
func (as *TodosService) Fetch(ctx context.Context, id string) (*Todo, *http.Response, error) {
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

	data := new(Todo)
	resp, err := as.client.Do(ctx, req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
