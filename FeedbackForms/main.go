package main

import (
	"html/template"
	"log"
	"net/http"
)

var temp = template.Must(template.ParseFiles("./template/static/index.html"))

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", renderTemplate)
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}


