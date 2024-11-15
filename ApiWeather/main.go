package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
)

type OpenWeatherResponse struct{
	Main struct{
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

type WeatherResponse struct{
	City string `json:"city"`
	Temperature float64 `json:"temperatura"`
}

func weatherHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Context-Type", "application/json")

	city := r.URL.Query().Get("city")
	if city ==""{
		city = "Porto Velho"
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == ""{
		http.Error(w, "API_KEY näo configura", http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil{
		http.Error(w, "Erro ao buscar dados da API", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		http.Error(w, "Error ao ler resposta", http.StatusInternalServerError)
		return
	}

	var openWeather OpenWeatherResponse
	if err := json.Unmarshal(body, &openWeather); err != nil {
		http.Error(w, "Erro ao decodificar Json", http.StatusInternalServerError)
		return
	}
	response := WeatherResponse{
		City: openWeather.Name,
		Temperature: openWeather.Main.Temp,
	}

	json.NewEncoder(w).Encode(response)
}

func main(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error ao carregar o arquivo .env")
	}

	http.HandleFunc("/weather", weatherHandler)

	log.Println("API rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}