package controllers

import (
	"bookstore-api/src/models"
	"bookstore-api/src/services"
	"encoding/json"
	"net/http"
)

// CreateBook handles the addition of a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	services.AddBook(book)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}
}

func GetBooks(w http.ResponseWriter) {
	books := services.GetBooks()
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		return
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	book, err := services.GetBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := services.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted"))
}
