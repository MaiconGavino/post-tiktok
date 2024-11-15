package main

import(
	"log"
	"net/http"
	"html/template"
	"github.com/joho/godotenv"
)

type OpenWeatherResponse struct{
	Main struct{
		Temp float64 `json:"temp"`
	}`json:"main"`
	Name string `json:"name"`
}

type WeatherResponse struct{
	City string `json:"city"`
	Temperature float64 `json:"temperature"`
}

func loadEnv(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file:", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	tmp := template.Must(template.ParseFiles("./template/index.html"))
	tmp.Execute(w, nil)
}


func main(){
	loadEnv()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)

	log.Println("Serve running on port 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))

}