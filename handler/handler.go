package handler

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Jika routing tidak sesuai
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello"))

}

func AboutHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("About Page"))
}

func ProfileHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Profile Page"))
}