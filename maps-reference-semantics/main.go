package main

import "fmt"

func addScore(scores map[string]int, name string, amount int) {
	scores[name] = scores[name] + amount
}

func incrementCount(m map[string]int, word string) {
	m[word] = m[word] + 1
}

func main() {
	counts := map[string]int{}

	incrementCount(counts, "apple")
	incrementCount(counts, "apple")
	incrementCount(counts, "banana")

	fmt.Println(counts)
}
