package main

import (
	"FeedbackForms/config"
	"FeedbackForms/handlers"
	"log"
	"net/http"
)

//func EnsureTableExists() {
//	query := `
//    CREATE TABLE IF NOT EXISTS feedbacks (
//        id SERIAL PRIMARY KEY,
//        name VARCHAR(100) NOT NULL,
//        email VARCHAR(100) NOT NULL,
//        rating INT NOT NULL,
//        comments TEXT,
//        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
//    )`
//	_, err := config.DB.Exec(query)
//	if err != nil {
//		log.Fatalf("Erro ao criar a tabela feedbacks: %v", err)
//	}
//	log.Println("Tabela feedbacks verificada/criada com sucesso.")
//}

func main() {

	config.ConnectDB()
	//EnsureTableExists()

	http.HandleFunc("/", handlers.RenderTemplate)
	http.HandleFunc("/submit", handlers.RegistreData)
	http.HandleFunc("/success", handlers.SuccessPage)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
