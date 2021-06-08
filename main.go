package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Email string `json:"Email"`
	Password string `json:"Password"`
	Bio string `json:"Bio"`
}

type Users []User

func baseUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go api")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Email: "cleiton@mail.go", Password: "myfirstgoapi123", Bio: "Hello, im a test user for this api"},
	}
	
	fmt.Println(w, "All users")
	json.NewEncoder(w).Encode(users)
}

func handeRequests() {
	http.HandleFunc("/", baseUrl)
	http.HandleFunc("/users", allUsers)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handeRequests()
}