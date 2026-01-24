package main

import "fmt"

func makeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func makeAdder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}

func main() {
	addTen := makeAdder(10)

	fmt.Println(addTen(5))  // 15
	fmt.Println(addTen(20)) // 30

	count := 0
	inc := func() int {
		count++
		return count
	}

	inc()
	inc()
}
