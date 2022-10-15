package main

import (
	"fmt"
	"github.com/paperlefthand/go-http-client/client"
)

func main() {
	status, err := client.GetStatusCode("https://pkg.go.dev/net/http")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(status)
}
