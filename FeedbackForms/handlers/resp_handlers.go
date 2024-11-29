package handlers

import (
	"FeedbackForms/config"
	"FeedbackForms/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseFiles("./template/index.html"))

func RenderTemplate(w http.ResponseWriter, r *http.Request) {
	temp.Execute(w, nil)
}

func RegistreData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	rating, err := strconv.ParseInt(r.FormValue("rating"), 10, 64)
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

	log.Printf("Inserindo no banco: Name=%s, Email=%s, Rating=%d, Comments=%s",
		resp.Name, resp.Email, resp.Rating, resp.Comments)

	query := "INSERT INTO feedbacks (name, email, rating, comments) VALUES ($1, $2, $3, $4)"

	_, err = config.DB.Exec(query, resp.Name, resp.Email, resp.Rating, resp.Comments)
	if err != nil {
		http.Error(w, "Erro ao salvar dados no banco", http.StatusInternalServerError)
		log.Printf("Erro ao salvar dados no banco: %v", err)
		return
	}

	http.Redirect(w, r, "/success", http.StatusSeeOther)
}

func SuccessPage(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "success.html", nil)
}
