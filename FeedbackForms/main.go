package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", renderTemplate)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
