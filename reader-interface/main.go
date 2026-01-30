package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func printContents(r io.Reader) {
	data, err := io.ReadAll(r)
	if err != nil {
		return
	}
	fmt.Println(string(data))
}

func main() {
	textReader := strings.NewReader("Hello from a string reader")
	printContents(textReader)

	file, err := os.Open("example.txt")
	if err != nil {
		return
	}
	defer file.Close()

	printContents(file)
}
