package Server

import (
	"fmt"
	"log"
	"net/http"
)

var Status = "off"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Start() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9012", nil))
}
