package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/skip2/go-qrcode"
)

type QrRequest struct {
	Text string `json:"text"`
}

func generateQRCode(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req QrRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Text == "" {
		http.Error(w, "Request Invalid", http.StatusBadRequest)
		return
	}

	qr, err := qrcode.Encode(req.Text, qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Error ao gerar QR Code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(qr)
}

func main() {
	fileServer := http.FileServer(http.Dir("./template")) 
	http.Handle("/", fileServer)
	http.HandleFunc("/generator", generateQRCode)
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}
