package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeEndPoint)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}

func HomeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
