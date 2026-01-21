package main

import "fmt"

func main() {
	var calculate func(int, int) int

	calculate = func(a int, b int) int {
		return a + b
	}

	fmt.Println(calculate(3, 4))
}
