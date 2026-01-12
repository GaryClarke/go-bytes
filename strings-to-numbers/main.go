package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "25"

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	number += 10

	result := strconv.Itoa(number)
	fmt.Println("Result:", result)
}
