package main

import "fmt"

func main() {
	// Challenge
	//Declare an array that holds three integers: 100, 200, and 300. Then print:
	//
	//The second number in the array
	//The entire array
	//Try doing this using both individual assignments and a shorthand declaration.

	numbers := [3]int{100, 200, 300}

	fmt.Println(numbers[1])
	fmt.Println(numbers)
}
