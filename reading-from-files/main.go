package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content))
}
