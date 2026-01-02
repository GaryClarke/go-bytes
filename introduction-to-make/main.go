package main

import "fmt"

func main() {
	scores := make(map[string]int)

	scores["Alice"] = 5
	scores["Bob"] = 3

	fmt.Println("Scores:", scores)
}
