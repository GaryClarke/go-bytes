package main

import (
	"errors"
	"fmt"
	"math"
)

func toSeconds(minutes int) (int, error) {
	if minutes < 0 {
		return 0, errors.New("minutes cannot be negative")
	}
	return minutes * 60, nil
}

func squareRoot(n int) (float64, error) {
	if n < 0 {
		return 0, errors.New("cannot take the square root of a negative number")
	}
	nFloat := float64(n)
	return math.Sqrt(nFloat), nil
}

func main() {
	sr, err := squareRoot(-4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", sr)
	}
}
