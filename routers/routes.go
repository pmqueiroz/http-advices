package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pmqueiroz/http-advices/status"
)

var port string = os.Getenv("PORT")
type User struct {
	Email 	string `json:"Email"`
	Password string `json:"Password"`
	Bio 		string `json:"Bio"`
}

type Users []User

func docs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/docs.html")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", docs).Methods("GET")
	router.HandleFunc("/advices", status.GetAllAdvices).Methods("GET")
	router.HandleFunc("/advices/{status}", status.GetAdviceByStatus).Methods("GET")

	log.Fatal(http.ListenAndServe(":" + port, router))
}

func Run() {
	message := "Running on port " + port
	
	fmt.Println(message)

	handleRequests()
}
