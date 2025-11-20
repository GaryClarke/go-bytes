package main

import "fmt"

func main() {
	// Challenge
	//Create a program that:
	//
	//Declares two int variables: a with value 10 and b with value 3
	//Performs integer division (a / b) and prints the result
	//Converts both a and b to float64
	//Performs float division with the converted values and prints the result
	//This demonstrates how type conversion lets you get fractional results from division.

	a := 10
	b := 3

	// Integer division (truncates)
	resultInt := a / b
	fmt.Println("Integer division:", resultInt) // 3

	// Convert to float64 for fractional result
	aFloat := float64(a)
	bFloat := float64(b)
	resultFloat := aFloat / bFloat
	fmt.Println("Float division:", resultFloat) // 3.3333333333333335

}
