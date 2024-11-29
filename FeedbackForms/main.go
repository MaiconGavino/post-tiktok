package main

import (
	"FeedbackForms/config"
	"FeedbackForms/handlers"
	"log"
	"net/http"
)

func main() {

	config.ConnectDB()

	http.HandleFunc("/", handlers.RenderTemplate)
	http.HandleFunc("/success", handlers.RegistreData)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
