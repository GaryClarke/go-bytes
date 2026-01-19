package main

import "fmt"

func main() {
	text := "Go ❤️"

	fmt.Println("Byte length:", len(text))
	for _, r := range text {
		fmt.Println(string(r))
	}
}
