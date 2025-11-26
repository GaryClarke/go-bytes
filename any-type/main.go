package main

import "fmt"

func describe(value any) {
	fmt.Println("Value:", value)
	fmt.Println("Type:", fmt.Sprintf("%T", value))
}

func logValue(label string, v any) {
	fmt.Println(label+":", v)
	fmt.Println("Type:", fmt.Sprintf("%T", v))
}

func main() {
	logValue("Age", 30)
	logValue("Name", "Alice")
	logValue("Price", 19.99)
}
