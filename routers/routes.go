package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/pmqueiroz/http-advices/advice"
)

type User struct {
	Email 	string `json:"Email"`
	Password string `json:"Password"`
	Bio 		string `json:"Bio"`
}

type Users []User

func docs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "docs/docs.html")
}

func handleRequests(port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", docs).Methods("GET")
	router.HandleFunc("/advices", advice.GetAllAdvices).Methods("GET")
	router.HandleFunc("/advices", advice.SuggestNewAdvice).Methods("POST")
	router.HandleFunc("/advices/{status}", advice.GetAdviceByStatus).Methods("GET")
	router.HandleFunc("/advices/search/{query}", advice.GetAdvicesByQuery).Methods("GET")

	log.Fatal(http.ListenAndServe(":" + port, router))
}

func Run(port *string) {
	alert := color.New(color.FgHiMagenta, color.Bold).PrintfFunc()

	message := "Running on port "

	fmt.Print(message)
	alert(*port)

	handleRequests(*port)
}