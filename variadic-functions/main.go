package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func greet(prefix string, names ...string) {
	for _, name := range names {
		fmt.Println(prefix, name)
	}
}

func multiplyAll(nums ...int) int {
	if len(nums) == 0 {
		return 1
	}
	result := 1
	for _, n := range nums {
		result *= n
	}
	return result
}

func main() {
	fmt.Println(multiplyAll(5))
	fmt.Println(multiplyAll(2, 3, 4))
	fmt.Println(multiplyAll())
}
