package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AzHAzizov/go_practice/2/SEMITRASH_API/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                    string
	bookResourcePrefix      string = apiPrefix + "/book"
	manyBooksResourcePrefix string = apiPrefix + "/books"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Could not find .env file")
	}

	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server ...")
	router := mux.NewRouter()

	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBooksResource(router, manyBooksResourcePrefix)

	log.Println("Router was initialized success")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
