package main

import (
	"fmt"
	"strings"
)

func main() {
	filename := "  report.txt  "

	trimmed := strings.TrimSpace(filename)
	valid := strings.HasSuffix(trimmed, ".txt")
	upper := strings.ToUpper(trimmed)

	fmt.Println("Trimmed:", trimmed)
	fmt.Println("Valid .txt:", valid)
	fmt.Println("Upper:", upper)
}
