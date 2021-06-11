package advice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var validStatusCode = [57]int64{100, 101, 102, 103, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308, 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 412, 413, 414, 415, 416, 417, 418, 421, 426, 428, 429, 431, 451, 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511}

type Advices struct {
	Advices []Advice `json:"advices"`
}

type Advice struct {
	Status 	int64 	  `json:"status"`
	Message  string  `json:"message"`
}

type SearchResult struct {
	Query    string	 `json:"query"`
	Total    int	    `json:"total_results"`
	Advices  []Advice  `json:"results"`
}

func fetchAdvices() Advices {
	file, err := os.Open("advices.json")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	
	byteValue, _ := ioutil.ReadAll(file)
	
	var data Advices

	json.Unmarshal(byteValue, &data)

	return data
}


func GetAdviceByStatus(w http.ResponseWriter, r *http.Request) {
	status := mux.Vars(r)["status"]
	statusId, _ := strconv.ParseInt(status, 10, 0)

	data := fetchAdvices()

	returnData := Advice{
		Status: 0,
		Message: "Invalid http status code. Check: https://developer.mozilla.org/docs/Web/HTTP/Status",
	}
	
	for _, sts := range validStatusCode {
		if sts == statusId {
			returnData = Advice{
				Status: 0,
				Message: "Status not implemented. Consider help this project on https://github.com/pmqueiroz/http-advices",
			}
		}
	}

	for i := 0; i < len(data.Advices); i++ {
		if data.Advices[i].Status == statusId {
			returnData = data.Advices[i]
		}
	}

	json.NewEncoder(w).Encode(returnData)
}

func GetAllAdvices(w http.ResponseWriter, r *http.Request){
	data := fetchAdvices()

	json.NewEncoder(w).Encode(data)
}

func GetAdvicesByQuery(w http.ResponseWriter, r *http.Request) {
	toWord := func(s string) string {return " " + s + " "}

	query := strings.ToLower(mux.Vars(r)["query"])
	
	data := fetchAdvices()

	var resultAdvices []Advice
	
	for _, advice := range data.Advices {
		if strings.Contains(strings.ToLower(advice.Message), toWord(query)) {
			resultAdvices = append(resultAdvices, advice)
		}
	}

	resultData := SearchResult{
		Query: query,
		Total: len(resultAdvices),
		Advices: resultAdvices,
	}

	json.NewEncoder(w).Encode(resultData)
}

func SuggestNewAdvice(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var data Advice

	err := json.Unmarshal(reqBody, &data)

	if err != nil {
		fmt.Println("error:", err)
	}

	json.NewEncoder(w).Encode(data)
}
