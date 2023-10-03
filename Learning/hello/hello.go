package main

import (
	"fmt"
	"os"

	"GO_Projects/Learning/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	fmt.Println(os.Getenv("GOPATH"))
}
