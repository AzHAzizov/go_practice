package main

import (
	"fmt"
	"log"
	"net/http"
)

// w ---> это куда пишем свой ответ
// r ---> это откуда берем запрос для ответа

func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi! I'm new web server!</h1>")
}

func RequestHander() {
	fmt.Fprintf()
	http.HandleFunc("/", GetGreet)
	log.Fatal(http.ListenAndServe(":8089", nil))
}

func main() {
	RequestHander()
}
