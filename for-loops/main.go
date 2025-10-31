package main

import "fmt"

func main() {
	// Challenge
	//Write a program that prints the following 3 lines using a for loop:
	//
	//Line 1
	//Line 2
	//Line 3
	//Use fmt.Println() and a loop counter.

	for i := 1; i <= 3; i++ {
		fmt.Println("Line", i)
	}
}
