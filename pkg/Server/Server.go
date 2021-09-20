package Server

import (
	"fmt"
	"log"
	"net/http"
)

var Status = "off"
var Temp = "0"

/*var pages = map[string]int{
	"rsc": 3711,
	"r":   2138,
	"gri": 1908,
}*/

func handlerDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server says " + Temp)
	fmt.Fprintf(w, Page, Temp)
}

func handlerLast(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server says " + Temp)
	fmt.Fprintf(w, "last %s", Temp)
	/*fmt.Fprintf(w, "Temp is  %s!", Temp)*/
}

func handlerHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server says " + Temp)
	fmt.Fprintf(w, "history %s", Temp)
	/*fmt.Fprintf(w, "Temp is  %s!", Temp)*/
}

func Start() {
	http.HandleFunc("/", handlerDefault)
	http.HandleFunc("/last", handlerLast)
	http.HandleFunc("/history", handlerHistory)
	log.Fatal(http.ListenAndServe(":9012", nil))
}
