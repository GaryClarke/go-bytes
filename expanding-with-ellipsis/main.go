package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func printAll(items ...string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	words := []string{"one", "two", "three"}

	printAll(words...)
}
