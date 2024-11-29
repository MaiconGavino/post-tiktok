package handlers

import (
	"FeedbackForms/config"
	"FeedbackForms/models"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/static/*.html"))

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		temp.Execute(w, nil)
	}
}

func RegistreData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	rating, err := strconv.Atoi(r.FormValue("rating"))
	if err != nil {
		http.Error(w, "Invalid rating", http.StatusBadRequest)
		return
	}
	r.ParseForm()
	resp := models.Resp{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Rating:   int64(rating),
		Comments: r.FormValue("comments"),
	}

	_, err = config.DB.Exec("INSERT INTO feedbacks (name, email, rating, comments) VALUES ($1, $2, $3, $4)",
		resp.Name, resp.Email, resp.Rating, resp.Comments)
	if err != nil {
		http.Error(w, "Error saving resp", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/sucess", http.StatusSeeOther)
}
