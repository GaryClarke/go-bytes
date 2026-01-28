package main

import (
	"io"
	"os"
)

func writeMessage(w io.Writer) {
	w.Write([]byte("Hello from io.Writer\n"))
}

func writeGreeting(w io.Writer) {
	w.Write([]byte("Hello from the writer interface\n"))
}

func main() {
	writeGreeting(os.Stdout)

	file, err := os.Create("greeting.txt")
	if err != nil {
		return
	}
	defer file.Close()
	writeGreeting(file)
}
