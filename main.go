package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     string  `json:"id"`
	Name   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func addBook(w http.ResponseWriter, r *http.Request) {

}

var books []Book

func main() {

	// Init Router
	r := mux.NewRouter()

	books = append(books, Book{"1", "111", "FirstBook", &Author{"Ping", "Chin"}})
	// Route Handler
	r.HandleFunc("/api/book", getAllBooks).Methods("GET")
	http.ListenAndServe(":9100", r)

}
