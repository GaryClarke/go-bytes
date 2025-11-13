package main

import "fmt"

func main() {
	// Challenge
	//Create a program that:
	//
	//Declares a score variable set to 100.
	//Creates a pointer to that variable.
	//Updates the value to 150 using the pointer.
	//Prints the updated value.
	score := 100
	ptr := &score
	*ptr = 150
	fmt.Println("Updated score:", score)
}
