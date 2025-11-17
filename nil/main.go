package main

import "fmt"

func main() {
	// Challenge
	// Try declaring the following without initializing them:
	//
	// A map[string]int
	// A []string slice
	// A int
	// Then write if statements to check if each one is nil, and print a message if it is.

	var scores map[string]int
	var names []string
	var value int

	if scores == nil {
		fmt.Println("scores map is nil")
	}

	if names == nil {
		fmt.Println("names slice is nil")
	}

	if value == nil {
		fmt.Println("value pointer is nil")
	}
}
