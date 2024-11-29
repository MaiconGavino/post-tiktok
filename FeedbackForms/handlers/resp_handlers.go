package handlers

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/static/*.html"))

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	temp.Execute(w, nil)
}
