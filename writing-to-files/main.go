package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("notes.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Lesson notes\n")
	file.WriteString("Writing files in Go\n")
	file.WriteString("Using os.Create\n")

	fmt.Println("notes.txt written successfully")
}
