package main

import "fmt"

func doubleFirst(values []int) {
	values[0] = values[0] * 2
}

func modify(values []int) {
	values[0] = 100
}

func addValue(values []int) []int {
	return append(values, 50)
}

func main() {
	nums := make([]int, 2, 2)
	nums[0], nums[1] = 10, 20

	nums2 := addValue(nums)
	fmt.Println("Original after append:", nums)
	fmt.Println("New slice after append:", nums2)
}
