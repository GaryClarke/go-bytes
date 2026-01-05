package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("Full:", now)
	fmt.Println("Date:", now.Format("2006-01-02"))
	fmt.Println("Time:", now.Format("15:04:05"))
}
