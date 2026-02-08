package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from the server")
	})

	http.ListenAndServe(":8080", nil)
}
