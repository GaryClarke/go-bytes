package main

import "fmt"

func divide(a, b int) (int, int, bool) {
	if b == 0 {
		return 0, 0, false
	}

	return a / b, a % b, true
}

func minMax(a, b int) (int, int, bool) {
	if a == b {
		return a, b, false
	}
	if a < b {
		return a, b, true
	}
	return b, a, true
}

func main() {
	min, max, ok := minMax(4, 4)
	if !ok {
		fmt.Println("The numbers are equal")
		return
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
