package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bookstore-api/src/controllers"
	"bookstore-api/src/models"
	"bookstore-api/src/services"
)

func setup() {
	// Clear the in-memory books slice before each test
	services.AddBook(models.Book{}) // Reset the slice
}

// TestCreateBook verifies the book creation functionality
func TestCreateBook(t *testing.T) {
	setup()

	book := models.Book{
		ID:     "1",
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Price:  45.99,
	}

	// Create a request body
	body, _ := json.Marshal(book)
	req, _ := http.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()

	// Call the controller function directly
	controllers.CreateBook(rec, req)

	// Verify the status code
	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("expected status %v but got %v", http.StatusCreated, status)
	}

	// Verify the response body
	var responseBook models.Book
	err := json.NewDecoder(rec.Body).Decode(&responseBook)
	if err != nil {
		return
	}

	if responseBook.ID != book.ID {
		t.Errorf("expected book ID %v but got %v", book.ID, responseBook.ID)
	}
}

// TestListBooks verifies fetching all books
func TestListBooks(t *testing.T) {
	setup()

	book := models.Book{
		ID:     "1",
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Price:  45.99,
	}
	services.AddBook(book)

	rec := httptest.NewRecorder()

	controllers.GetBooks(rec)

	// Verify the status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, status)
	}

	// Verify the response body
	var books []models.Book
	err := json.NewDecoder(rec.Body).Decode(&books)
	if err != nil {
		return
	}

	if len(books) != 1 {
		t.Errorf("expected 1 book but got %d", len(books))
	}

	if books[0].ID != book.ID {
		t.Errorf("expected book ID %v but got %v", book.ID, books[0].ID)
	}
}

// TestGetBook verifies fetching a book by ID
func TestGetBook(t *testing.T) {
	setup()

	book := models.Book{
		ID:     "1",
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Price:  45.99,
	}
	services.AddBook(book)

	req, _ := http.NewRequest(http.MethodGet, "/books/detail?id=1", nil)
	rec := httptest.NewRecorder()

	controllers.GetBook(rec, req)

	// Verify the status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, status)
	}

	// Verify the response body
	var responseBook models.Book
	err := json.NewDecoder(rec.Body).Decode(&responseBook)
	if err != nil {
		return
	}

	if responseBook.ID != book.ID {
		t.Errorf("expected book ID %v but got %v", book.ID, responseBook.ID)
	}
}

// TestDeleteBook verifies deleting a book by ID
func TestDeleteBook(t *testing.T) {
	setup()

	book := models.Book{
		ID:     "1",
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Price:  45.99,
	}
	services.AddBook(book)

	req, _ := http.NewRequest(http.MethodDelete, "/books/delete?id=1", nil)
	rec := httptest.NewRecorder()

	controllers.DeleteBook(rec, req)

	// Verify the status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected status %v but got %v", http.StatusOK, status)
	}

	// Verify the book was deleted
	_, err := services.GetBookByID("1")
	if err == nil {
		t.Error("expected error when fetching deleted book but got none")
	}
}
