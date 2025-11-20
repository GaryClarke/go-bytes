package main

import "fmt"

func main() {
	// Challenge
	//Create a program that:
	//
	//Declares a float32 variable for a price
	//Declares a float64 variable for a temperature
	//Uses fmt.Printf with %T to print the type of both variables
	//Declares a third variable using := with a decimal value and prints its type

	var price float32 = 29.99
	var temp float64 = 25.5

	fmt.Printf("Price type: %T\n", price)
	fmt.Printf("Temperature type: %T\n", temp)

	height := 1.75
	fmt.Printf("Height type: %T\n", height)

}
