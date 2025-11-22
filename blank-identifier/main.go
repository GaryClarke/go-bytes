package main

import "fmt"

func getUserInfo() (string, int) {
	return "Alice", 35
}

func main() {
	// Challenge
	//Create a function that returns two values: a string and an integer.
	//
	//Call the function and:
	//
	//Use the string, ignore the integer.
	name, _ := getUserInfo()
	fmt.Println(name)
	//
	//Then, ignore the string and use the integer.
	_, age := getUserInfo()
	fmt.Println(age)

}
