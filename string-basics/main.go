package main

import "fmt"

func main() {
	// Challenge
	//Write a program that:
	//
	//Declares a string variable for your first name.
	//Declares a string variable for your favorite programming language.
	//Prints a message like: My name is Alice and I love Go.
	//(Use string concatenation - not fmt.Sprintf yet.)

	name := "Alice"
	name = "Gary"
	language := "Go"
	fmt.Println("My name is " + name + " and I love " + language + ".")
}
