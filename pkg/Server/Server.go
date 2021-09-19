package Server

import (
	"fmt"
	"log"
	"net/http"
)

var Status = "off"

var Temp = "0"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server says " + Temp)
	fmt.Fprintf(w, Page, Temp)
	/*fmt.Fprintf(w, "Temp is  %s!", Temp)*/

}

func Start() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9012", nil))
}
