package main

import (
	"example/pkg/env"
	"fmt"
)

func main() {
	// Load the environment settings
	if err := env.Load("local", true); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
