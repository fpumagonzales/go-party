package main

import (
	"fmt"

	"github.com/fpumagonzales/go-party/pkg/party"
)

func main() {
	fmt.Println("Hello Party!")

	jennifer := party.NewParty("Jennifer", "White", "jwhite@go.com")

	fmt.Printf("%+v", jennifer)
}
