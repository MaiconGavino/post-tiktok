package handlers

import (
	"forms/config"
	"forms/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func FormsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	user := models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error generating password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	_, err = config.DB.Exec("INSERT INTO users (firstname, lastname, email, password) VALUES ($1, $2, $3, $4)",
		user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/sucess", http.StatusSeeOther)
}

func SucessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Registro realizado com sucesso!"))
}
