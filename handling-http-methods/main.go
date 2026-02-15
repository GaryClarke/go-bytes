package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintln(w, "Submit page")
		case http.MethodPost:
			fmt.Fprintln(w, "Submission received")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
