package main

import "fmt"

func sayHello() {
	fmt.Println("Hello from a goroutine")
}

func main() {
	go sayHello()
	fmt.Println("Hello from main")
}
