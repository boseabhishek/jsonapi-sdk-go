package data

import (
	"fmt"
	"go-rest/thirdparty"
	"net/http"
)

type Data struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func NewData() *Data {
	return &Data{}
}

func (d *Data) Fetch(id string) (*Data, *http.Response, error) {
	// create the resource
	u := fmt.Sprintf("posts/%s", id)

	// create anew thirdparty client
	client := thirdparty.NewClient()

	// create a http request
	req, err := client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	data := new(Data)
	resp, err := client.Do(req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
