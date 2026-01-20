package main

import "fmt"

func greet() {
	fmt.Println("Hello!")
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func main() {
	var op func(int, int) int

	op = multiply
	fmt.Println(op(10, 2))

	op = divide
	fmt.Println(op(10, 2))
}
