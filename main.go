package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server running on Port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
