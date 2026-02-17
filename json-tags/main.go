package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	product := Product{
		ID:   1,
		Name: "Notebook",
	}
	data, _ := json.Marshal(product)

	fmt.Println(string(data))
}
