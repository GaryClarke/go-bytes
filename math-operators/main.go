package main

import (
	"fmt"
	"math"
)

func main() {
	// Uses / Use Cases
	//Doing calculations in business logic
	//Processing user input or numerical data
	//Using math constants like Pi in geometry-related code
	//Performing square root or power operations

	// Basic arithmetic
	a := 10
	b := 3

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b) // integer division
	fmt.Println("a % b =", a%b) // modulus (remainder)

	// Float division
	x := 10.0
	y := 3.0
	fmt.Println("x / y =", x/y)

	// Using parentheses for order of operations
	c := 12.0
	d := 5.0
	e := 2.0
	result := c + (d*e)*d
	fmt.Println("Result using parentheses:", result)

	// Using math package
	fmt.Println("Pi:", math.Pi)
	fmt.Println("Square root of 16:", math.Sqrt(16))
}
