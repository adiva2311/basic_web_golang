package main

import (
	"basic-web/handler"
	"log"
	"net/http"
)

func main() {
	// Membuat server baru
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/about", handler.AboutHandler)
	mux.HandleFunc("/profile", handler.ProfileHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("Starting on Port 9000")

	// Menjalankan Server
	err := http.ListenAndServe("127.0.0.1:9000", mux)
	log.Fatal(err)
}