package main

import "fmt"

func main() {
	// Write a program that checks a score and prints a message based on this logic:

	// If the score is 90 or above, print "Grade: A"
	// Else if the score is 80 or above, print "Grade: B"
	// Else if the score is 70 or above, print "Grade: C"
	// Else, print "Grade: F"

	// Use an int variable named score

	score := 75
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade F")
	}
}
