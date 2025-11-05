package main

import "fmt"

func main() {
	// Challenge
	//Create a program with three variables:
	//
	//loggedIn (true or false)
	//isStaff (true or false)
	//hasPermission (true or false)
	//Write a condition that prints "Action allowed" only if:
	//
	//The user is logged in
	//AND they are staff or have permission
	//Then test different combinations of values.

	loggedIn := true
	isStaff := false
	hasPermission := false

	if loggedIn && (isStaff || hasPermission) {
		fmt.Println("Action allowed")
	} else {
		fmt.Println("Action denied")
	}
}
