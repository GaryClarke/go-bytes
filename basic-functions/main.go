package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello", name)
}

func square(n int) int {
	return n * n
}

func double(n int) int {
	return n * 2
}

func main() {
	// Challenge
	//Write a function named double that takes an integer and returns twice its value.
	//
	//In your main() function, call double(4) and print the result using fmt.Println().

	result := double(4)
	fmt.Println(result)
}
