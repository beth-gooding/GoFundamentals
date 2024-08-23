package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {
	id := uuid.New()
	fmt.Printf("Generate UUID: %s\n", id)

	quote := quote.Glass()
	fmt.Println("Random Quote:", quote)
}
