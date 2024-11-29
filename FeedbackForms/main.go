package main

import (
	"FeedbackForms/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.RenderTemplate)
	http.HandleFunc("/sucess", handlers.RegistreData)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
