package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Nome     string `json:"nome"`
	Tel      string `json:"tel"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

var users = []User{}

func createUser( w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, "Error ao processar os dados", http.StatusBadRequest)
		return
	}

	if user.Nome == "" || user.Tel == "" || user.Email == "" || user.Password == ""{
		http.Error(w, "Todos os campus precisam ser preenchidos", http.StatusBadRequest)
		return
	}
	users = append(users, user)
	log.Println("Novo usuário adicionado: %+v/n", user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
 }

func main() {
	fileServe := http.FileServer(http.Dir("./template/"))
	http.Handle("/", fileServe)
	http.HandleFunc("/create", createUser)
	fmt.Println("Server run in port 8080")
	http.ListenAndServe(":8080", nil)

}
