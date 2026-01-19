package main

import "fmt"

type UserID int

type ProductID int

func printProduct(id ProductID) {
	fmt.Println("Product ID:", id)
}

func main() {
	var rawID int = 42
	printProduct(ProductID(rawID))
}
