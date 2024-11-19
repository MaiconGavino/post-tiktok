package main

import (
	"forms/config"
	"forms/handlers"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	// Configura para servir arquivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rota para renderizar o formulário
	http.HandleFunc("/", handlers.RegistrationFormHandler)

	// Rota para registrar o usuário
	http.HandleFunc("/register", handlers.RegisterUserHandler)

	// Rota para a página de sucesso
	http.HandleFunc("/success", handlers.SuccessHandler)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
