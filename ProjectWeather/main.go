package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OpenWeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

type WeatherResponse struct {
	City        string
	Temperature float64
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		city = "Porto Velho" // Cidade padrão
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		http.Error(w, "API_KEY not configured", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data from API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	var openWeather OpenWeatherResponse
	if err := json.Unmarshal(body, &openWeather); err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	weatherData := WeatherResponse{
		City:        openWeather.Name,
		Temperature: openWeather.Main.Temp,
	}

	// Renderiza o template com os dados dinâmicos
	tmpl := template.Must(template.ParseFiles("./template/index.html"))
	tmpl.Execute(w, weatherData)
}

func main() {
	loadEnv()

	fs := http.FileServer(http.Dir("./template/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/weather", weatherHandler)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

