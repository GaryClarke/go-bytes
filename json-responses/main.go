package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		product := Product{
			ID:    10,
			Name:  "Notebook",
			Price: 12.99,
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(product)
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
