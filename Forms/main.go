package main

import (
	"forms/config"
	"forms/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	defer config.DB.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.FormsHandler).Methods("GET")
	r.HandleFunc("/register", handlers.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/sucess", handlers.SucessHandler).Methods("GET")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
