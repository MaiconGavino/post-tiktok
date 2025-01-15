package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TextRequest struct{
	Text string `json:"text"`
}

func processText(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var text TextRequest
	err:= json.NewDecoder(r.Body).Decode(&text)
	if err != nil || text.Text == ""{
		http.Error(w, "Request Invalido", http.StatusBadRequest)
		return
	}
	log.Println(text)
}

func main(){
	fileServe := http.FileServer(http.Dir("./template"))
	http.Handle("/", fileServe)
	http.HandleFunc("/convert", processText)
	fmt.Println("Iniciando o servido na porta 8080")
	http.ListenAndServe(":8080", nil)
}