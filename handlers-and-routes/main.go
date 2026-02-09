package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home page")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Contact us")
	})

	http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Help page")
	})

	http.ListenAndServe(":8080", nil)
}
