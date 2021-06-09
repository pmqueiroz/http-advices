package status

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Advices struct {
	Advices []Advice `json:"advices"`
}

type Advice struct {
	Status 	int64 	  `json:"status"`
	Message  string  `json:"message"`
}


func SendAdvice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	statusId, _ := strconv.ParseInt(vars["status"], 10, 0)

	file, err := os.Open("advices.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	
	byteValue, _ := ioutil.ReadAll(file)
	
	var data Advices

	json.Unmarshal(byteValue, &data)

	returnData := Advice{
		Status: 0,
		Message: "Status not implemented. Consider help this project on https://github.com/pmqueiroz/http-advices",
	}

	for i := 0; i < len(data.Advices); i++ {
		if data.Advices[i].Status == statusId {
			returnData = data.Advices[i]
		}
	}

	json.NewEncoder(w).Encode(returnData)
}

func RequiredStatus(w http.ResponseWriter, r *http.Request){

}
