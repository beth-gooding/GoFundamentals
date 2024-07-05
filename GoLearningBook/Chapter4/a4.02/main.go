package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	return map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
}

func getUser(id string) (string, bool) {
	users := getUsers()

	name, exists := users[id]

	return name, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("No user id passed in")
		// Add an exit statement here
		os.Exit(1)
	}

	name, exists := getUser(os.Args[1])

	if !exists {
		fmt.Printf("No user with id (%v)\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("Hi,", name)
}
