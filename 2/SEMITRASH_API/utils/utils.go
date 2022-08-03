package utils

import (
	"github.com/AzHAzizov/go_practice/2/SEMITRASH_API/hendlers"
	"github.com/gorilla/mux"
)

func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", hendlers.GetBookById).Methods("GET")
	router.HandleFunc(prefix, hendlers.CreateBook).Methods("POST")
	router.HandleFunc(prefix+"/{id}", hendlers.UpdateBook).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", hendlers.DeleteBook).Methods("DELETE")
}

func BuildManyBooksResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix, hendlers.GetAllBooks).Methods("GET")
}
