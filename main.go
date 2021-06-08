package main

import (
	"fmt"
	"log"
	"net/http"
)

func baseUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go api")
}


func handeRequests() {
	http.HandleFunc("/", baseUrl)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handeRequests()
}