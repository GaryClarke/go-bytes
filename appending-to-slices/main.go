package main

import "fmt"

func main() {
	values := make([]int, 0, 2)

	for i := 1; i <= 5; i++ {
		values = append(values, i)
		fmt.Println("values:", values, "len:", len(values), "cap:", cap(values))
	}
}
