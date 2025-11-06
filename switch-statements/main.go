package main

import "fmt"

func main() {
	// Challenge
	//Create a program that checks the value of a status variable and prints a message:
	//
	//If status is "success", print "Operation completed successfully."
	//If status is "error", print "Something went wrong."
	//If status is "loading", print "Please wait..."
	//For anything else, print "Unknown status."
	//Use a switch statement to handle the logic.

	status := "terminated"

	switch status {
	case "success":
		fmt.Println("Operation completed successfully.")
	case "error":
		fmt.Println("Something went wrong.")
	case "loading":
		fmt.Println("Please wait...")
	default:
		fmt.Println("Unknown status")
	}
}
