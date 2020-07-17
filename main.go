package main

import (
	"fmt"
	"go-rest/data"
)

func main() {

	d, _, err := data.NewData().Fetch("1")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}

	fmt.Println(d)

}
