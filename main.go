package main

import (
	"fmt"
	"go-rest/account"
)

func main() {

	client := account.NewClient()
	d, _, err := client.Accounts.Fetch("1")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}

	fmt.Println(d)

}
