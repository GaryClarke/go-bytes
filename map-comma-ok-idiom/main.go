package main

import "fmt"

func main() {
	emails := map[string]string{
		"alice": "alice@example.com",
		"bob":   "bob@example.com",
	}

	username := "sarah"

	email, ok := emails[username]
	if !ok {
		fmt.Println("User not found:", username)
		return
	}

	fmt.Println("Email:", email)
}
