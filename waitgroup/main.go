package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	var wg sync.WaitGroup

	for _, n := range numbers {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(n)
	}

	wg.Wait()
}
