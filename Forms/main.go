package main

import (
	"forms/config"
	"forms/handlers"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.RegistrationFormHandler)
	http.HandleFunc("/register", handlers.RegisterUserHandler)
	http.HandleFunc("/success", handlers.SuccessHandler)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
