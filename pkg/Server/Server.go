package Server

import (
	"XU4PowerManager/pkg/HWReader"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Status = "off"

/*var pages = map[string]int{
	"rsc": 3711,
	"r":   2138,
	"gri": 1908,
}*/

func handlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, Page, 12)
}

func handlerLast(w http.ResponseWriter, r *http.Request) {
	//jsonString, _ := json.Marshal(HWReader.GetLast())
	//fmt.Fprintf(w, string(jsonString))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HWReader.GetLast())

}

func handlerHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HWReader.GetHistory())
}

func Start() {
	http.Handle("/", http.FileServer(http.Dir("./html")))
	http.HandleFunc("/last", handlerLast)
	http.HandleFunc("/history", handlerHistory)
	log.Fatal(http.ListenAndServe(":9012", nil))
}
