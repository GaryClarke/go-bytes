package main

import "fmt"

func main() {
	// Challenge
	//Create a slice of 3 integers and print each of the following:
	numbers := []int{10, 20, 30}
	//
	//The full slice
	fmt.Println("Original slice:", numbers)
	//The first value
	fmt.Println("First item:", numbers[0])
	//The length of the slice
	fmt.Println("Original slice length:", len(numbers))
	//Then, append a fourth number to the slice and print the result.
	numbers = append(numbers, 40, 50, 60)
	//
	fmt.Println("New slice length:", len(numbers))
	//You can use len() to get the length of a slice.
	fmt.Println("New slice:", numbers)
}
