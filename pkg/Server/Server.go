package Server

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"XU4PowerManager/pkg/HWReader"
)

// content holds our static web server content.
//go:embed html/live.html
var content embed.FS
var Status = "off"

func handleDefault(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./html/live.html")	//Serve a file
	//fmt.Fprintf(w, "Page, 12")				//Serve a string
	http.Redirect(w, r, "/live", 303) //Redirect
}
func serveHistory(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./html/history.html")
	b, _ := content.ReadFile("html/history.html")
	w.Write(b)
}

func serveLive(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, "./html/live.html")
	b, _ := content.ReadFile("html/live.html")
	w.Write(b)
}

func serveLast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HWReader.GetSystemStatus())
}
func serveLog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HWReader.GetHistory())
}
func serveLog2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(HWReader.GetLog())
}

func Start() {
	Status = "on"
	//http.Handle("/", http.FileServer(http.Dir("./html")))
	http.HandleFunc("/", handleDefault)
	http.HandleFunc("/live", serveLive)
	http.HandleFunc("/history", serveHistory)
	http.HandleFunc("/last", serveLast)
	http.HandleFunc("/log", serveLog)
	http.HandleFunc("/log2", serveLog2)
	log.Fatal(http.ListenAndServe(":9012", nil))
	fmt.Println("Server: " + Status)
}
