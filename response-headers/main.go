package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/created", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Some content")
	})
	http.ListenAndServe(":8080", nil)
}
