package main

import (
	"context"
	"fmt"
	"go-rest/jsonapi"
)

func main() {

	client := jsonapi.NewClient()
	//d, _, err := client.Accounts.Fetch(context.Background(), "79c9d274-d72f-11ea-87d0-0242ac130003")
	//acs, _, err := client.Accounts.List(context.Background())
	_, _, err := client.Accounts.Create(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}

	/* for i, ac := range acs {
		fmt.Printf("%v. %v\n", i+1, ac)
	} */

	fmt.Printf("------%+v", err)

}
