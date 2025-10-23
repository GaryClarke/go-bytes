package main

import "fmt"

func main() {
	// Write a Go program that creates a simple user profile. Declare variables for:
	//
	//1. A `string` for your favorite programming language
	//2. An `int` for how many years you've been programming
	//3. A `bool` for whether you prefer coding in the morning or evening (true for morning, false for evening)
	//4. A `float64` for your programming confidence level (0.0 to 1.0)

	favoriteLanguage := "Go"
	yearsProgramming := 1
	var morningCoder bool = true
	confidenceLevel := 0.7

	fmt.Println("Favorite Language:", favoriteLanguage)
	fmt.Println("Years Programming:", yearsProgramming)
	fmt.Println("Morning Coder:", morningCoder)
	fmt.Println("Confidence Level:", confidenceLevel)
}
