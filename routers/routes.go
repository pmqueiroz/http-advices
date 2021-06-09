package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pmqueiroz/http-advices/status"
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

func adviceDoc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/advice.html")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", baseUrl)
	router.HandleFunc("/advice", adviceDoc)
	router.HandleFunc("/advice/{status}", status.SendAdvice)
	log.Fatal(http.ListenAndServe(port, router))
}

func Run() {
	message := "Running on port " + port
	
	fmt.Println(message)

	handleRequests()
}
