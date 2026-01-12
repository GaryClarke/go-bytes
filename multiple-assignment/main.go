package main

import "fmt"

func getPoint() (int, int) {
	return 3, 7
}

func main() {
	a := 3
	b := 8
	fmt.Println("Before swap:", a, b)

	a, b = b, a
	fmt.Println("After swap:", a, b)
}
