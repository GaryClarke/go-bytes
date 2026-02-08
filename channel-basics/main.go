package main

import "fmt"

func main() {
	// Create a channel for integers
	numbers := make(chan int)

	// Start a goroutine that sends a number
	go func() {
		// Send 100 into the channel
		numbers <- 100
	}()

	// Receive the number from the channel
	value := <-numbers
	fmt.Println("Received:", value)

}
