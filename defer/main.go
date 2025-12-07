package main

import "fmt"

func process() {
	fmt.Println("Starting work")
	defer fmt.Println("Cleaning up")
	fmt.Println("Doing work")
}

func demo() {
	x := 10
	defer fmt.Println("Deferred value:", x)
	x = 20
	fmt.Println("Updated value:", x)
}

func show() {
	for i := 1; i <= 3; i++ {
		defer fmt.Println("Deferred:", i)
	}
	fmt.Println("End of function")
}

func saveReport() {
	fmt.Println("Saving report")
	defer fmt.Println("Closing resources")
	defer fmt.Println("Cleaning temporary data")
	fmt.Println("Writing data")
}

func main() {
	saveReport()
}
