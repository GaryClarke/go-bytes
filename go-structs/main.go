package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	// Challenge
	//Create a struct called Book with the following fields:
	//
	//Title (string)
	//Author (string)
	//Year (int)
	//Then initialize a Book, set the fields to anything you like, and print them out using fmt.Println().
	b := Book{
		Title:  "Go in Practice",
		Author: "Alice Smith",
		Year:   2025,
	}

	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Year:", b.Year)
}
