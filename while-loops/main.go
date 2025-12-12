package main

import "fmt"

func main() {
	n := 1

	for n <= 100 {
		fmt.Println("n is:", n)
		n = n * 2
	}
}
