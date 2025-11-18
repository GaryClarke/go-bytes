package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	// Challenge
	//Declare variables of the following types without assigning a value:
	//
	// int
	// string
	// bool
	// []float64
	//
	//A struct of your choice
	//
	//Print each one to confirm their zero values.
	var u User

	fmt.Println(u)
}
