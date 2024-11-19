package handlers

import (
	"forms/config"
	"forms/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("./template/index.html"))

// Renderiza o formulário de registro
func RegistrationFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

// Processa o registro do usuário
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	user := models.User{
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	_, err = config.DB.Exec("INSERT INTO users (firstname, lastname, email, password) VALUES ($1, $2, $3, $4)",
		user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, "Error saving user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/success", http.StatusSeeOther)
}

// Página de sucesso
func SuccessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Registro realizado com sucesso!"))
}
