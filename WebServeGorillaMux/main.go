package main

import(
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"sync"
	"github.com/gorilla/mux"
)

type Item struct{
	ID string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{}
var mu sync.Mutex

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/api/items", GetItemsHandler).Methods("GET")
	r.HandleFunc("/api/items", CreateItemHandler).Methods("POST")

	fmt.Println("Start Serve Running on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the Golang with Gorilla Mux Web Serve Project"))
}

func GetItemsHandler( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(items)
}
func CreateItemHandler(w http.ResponseWriter, r *http.Request){
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	newItem.ID = fmt.Sprintf("%d", len(items)+1)
	items = append(items, newItem )

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)

}