package main

import "fmt"

func describe(value any) {
	if s, ok := value.(string); ok {
		fmt.Println("String:", s)
		return
	}

	if i, ok := value.(int); ok {
		fmt.Println("Int:", i)
		return
	}

	fmt.Printf("Unknown type: %T\n", value)
}

func main() {
	describe("hello")
	describe(42)
	describe(true)
}
