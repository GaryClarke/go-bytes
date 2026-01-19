package main

import "fmt"

func main() {
	numbers := []int{5, 10, 15}
	copied := make([]int, len(numbers))

	copy(copied, numbers)

	copied[2] = 100

	fmt.Println("numbers:", numbers)
	fmt.Println("copied:", copied)
}
