package main

import "fmt"

func square(n int) (result int) {
	result = n * n
	return
}

func divide(a, b float64) (quotient float64, ok bool) {
	if b == 0 {
		ok = false
		return
	}

	quotient = a / b
	ok = true
	return
}

func stats(nums []int) (total int, count int) {
	for _, n := range nums {
		total += n
		count++
	}
	return
}

func main() {
	total, count := stats([]int{3, 5, 7})
	fmt.Println("Total:", total)
	fmt.Println("Count:", count)
}
