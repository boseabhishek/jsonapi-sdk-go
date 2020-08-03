package main

import (
	"context"
	"fmt"
	"go-rest/jsonapi"
)

func main() {

	client := jsonapi.NewClient()
	//d, _, err := client.Accounts.Fetch(context.Background(), "1")
	acs, _, err := client.Accounts.List(context.Background())
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}

	for i, ac := range acs {
		fmt.Printf("%v. %v\n", i+1, ac)
	}

}
