package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/profile", profileHandler)

	log.Println("Starting on Port 9000")

	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("About Page"))
}

func profileHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Profile Page"))
}