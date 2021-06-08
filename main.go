package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port string = ":8081"

type User struct {
	Email 	string `json:"Email"`
	Password string `json:"Password"`
	Bio 		string `json:"Bio"`
}

type Users []User

func baseUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Go api")
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Email: "cleiton@mail.go", Password: "myfirstgoapi123", Bio: "Hello, im a test user for this api"},
	}
	
	json.NewEncoder(w).Encode(users)
}

func handeRequests() {
	http.HandleFunc("/", baseUrl)
	http.HandleFunc("/users", allUsers)
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	message := "Running on port " + port

	fmt.Println(message)
	handeRequests()
}