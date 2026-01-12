package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name and age: ")
	fmt.Scan(&name, &age)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}
