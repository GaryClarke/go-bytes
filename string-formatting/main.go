package main

import "fmt"

func main() {
	// Challenge
	//Create variables for a product:
	//
	//Name: "Notebook"
	//Price: 9.99
	//Use fmt.Printf to print: Notebook costs $9.99
	//
	//Use fmt.Sprintf to build the same string and store it in a variable.
	//
	//Print the result using fmt.Println.
	name := "Notebook"
	price := 9.99
	message := fmt.Sprintf("%s costs %.2f\n", name, price)
	fmt.Println(message)

}
