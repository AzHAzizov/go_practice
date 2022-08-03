package hendlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AzHAzizov/go_practice/2/SEMITRASH_API/models"
	"github.com/gorilla/mux"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error while parsing happen:", err)
		msg := models.Message{Message: "do not use param ID as uncased to int type"}
		json.NewEncoder(w).Encode(msg)
		w.WriteHeader(400)
		return
	}

	book, ok := models.FindBookById(id)
	if !ok {
		log.Println("Book not found :", id)
		msg := models.Message{Message: fmt.Sprintf("Book by id %d is not found", id)}
		json.NewEncoder(w).Encode(msg)
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(book)
	return

}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	log.Println("Start create new book ....")
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)

	fmt.Println(book)

	// json.NewDecoder( вычитываем откуда ).Decode( вычитываем куда )
	if err != nil {

		log.Println(err)
		msg := models.Message{Message: "Invalid type of data to create new BOOK"}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}

	var newBookId int = len(models.DB) + 1
	book.ID = newBookId
	models.DB = append(models.DB, book)

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	log.Println("Start update book .....")

	// Copy from get book by id
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error while parsing happen:", err)
		msg := models.Message{Message: "do not use param ID as uncased to int type"}
		json.NewEncoder(w).Encode(msg)
		w.WriteHeader(400)
		return
	}

	key, ok := models.FindBookKey(id)

	if !ok {
		log.Println("Book not found to update:", id)
		msg := models.Message{Message: fmt.Sprintf("Book by id %d is not found", id)}
		json.NewEncoder(w).Encode(msg)
		w.WriteHeader(404)
		return
	}

	var newBook models.Book
	json.NewDecoder(r.Body).Decode(&newBook)

	fmt.Println(models.DB[key])

	models.DB[key] = newBook

	log.Println(models.DB)

	msg := models.Message{Message: "Book was updated success"}
	w.WriteHeader(200) // ?????????????????
	json.NewEncoder(w).Encode(msg)
	return

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	log.Println("Start remove book:", id)
	if err != nil {
		log.Println("error while parsing happen:", err)
		msg := models.Message{Message: "do not use param ID as uncased to int type"}
		json.NewEncoder(w).Encode(msg)
		w.WriteHeader(400)
		return
	}

	_, ok := models.FindBookById(id)
	if !ok {
		log.Println("Book not found to remove:", id)
		msg := models.Message{Message: fmt.Sprintf("Book by id %d is not found", id)}
		json.NewEncoder(w).Encode(msg)
		w.WriteHeader(404)
		return
	}

	key, ok := models.FindBookKey(id)
	models.RemoveFromBooks(key)

	// TODO :: remove BOOK

	msg := models.Message{Message: "Book was removed success"}
	w.WriteHeader(200) // ?????????????????
	json.NewEncoder(w).Encode(msg)
	return
}
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	initHeaders(w)
	log.Println("Get all books from DB")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.DB)
}
