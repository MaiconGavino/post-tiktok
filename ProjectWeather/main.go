package main

import(
	"log"
	"os"
	"net/http"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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


func weatherHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	city := r.URL.Query().Get("City")
	if city == ""{
		city = "Campinas"
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == ""{
		http.Error(w, "Error fetching data from API", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data from API", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}
	var openWeather OpenWeatherResponse
	if err := json.Unmarshal(body, &openWeather); err != nil{
		http.Error(w, "Error decoding Json", http.StatusInternalServerError)
		return
	}
	response := WeatherResponse{
		City: openWeather.Name,
		Temperature : openWeather.Main.Temp,
	}
	json.NewEncoder(w).Encode(response)
}



func main(){
	loadEnv()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("weather", weatherHandler)

	log.Println("Serve running on port 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))

}