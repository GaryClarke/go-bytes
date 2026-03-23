package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	doWork()

	fmt.Println("Program continues")
}

func doWork() {
	panic("unexpected error")
}
