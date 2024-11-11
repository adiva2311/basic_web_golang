package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", homeHandler)

	log.Println("Starting on Port 9000")

	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}