package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// Mengambil Query String
	id := r.URL.Query().Get("id")

	// Membuat validasi query string menjadi number
	idNum, err := strconv.Atoi(id)
	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	// Untuk menampilkan query yang sudah diambil
	fmt.Fprintf(w, "Product Page: %d", idNum)
}