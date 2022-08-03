package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	port string = "8080"
	db   []Pizza
)

func init() {
	pizza1 := Pizza{
		ID:       1,
		Diameter: 22,
		Price:    500.50,
		Title:    "Pepperoni",
	}

	pizza2 := Pizza{
		ID:       2,
		Diameter: 25,
		Price:    650.23,
		Title:    "BBQ",
	}
	pizza3 := Pizza{
		ID:       3,
		Diameter: 22,
		Price:    450,
		Title:    "Margaritta",
	}

	db = append(db, pizza1, pizza2, pizza3)
}

type Pizza struct {
	ID       int     `json:"id"`
	Diameter int     `json:"diameter"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
}

func FindPizzaById(id int) (Pizza, bool) {
	for _, pizza := range db {
		if pizza.ID == id {
			return pizza, true
		}
	}

	return Pizza{}, false
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func GetAllPizzas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(db)

}

func GetPizzaById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := ErrorMessage{Message: fmt.Sprintf("Could not convert to id data: %v", vars["id"])}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}

	log.Println("Trying to get pizza with id", id)
	pizza, ok := FindPizzaById(id)

	if ok {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(pizza)
		return
	}
	ms := ErrorMessage{Message: fmt.Sprintf("Pizza wih id %d not found", id)}
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(ms)

}

func main() {
	log.Println("Pizza application is works correct")

	// init router
	router := mux.NewRouter()                                     // Инициализируем маршрутизатор gorilla mux
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")     // определяем для него первый handler который обрабатывает запрос для получение всех ПИЦ
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET") // второый handler для того что получть пицу по id
	log.Fatal(http.ListenAndServe(":"+port, router))              // устанавливаем наш роутор для данного listener а
}
