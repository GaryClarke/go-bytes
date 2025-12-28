package main

import (
	"fmt"
)

func loadConfig() error {
	return nil
}

func getUsername() (string, error) {
	return "Gary", nil
}

func main() {
	if name, err := getUsername(); err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Username:", name)
	}
}
