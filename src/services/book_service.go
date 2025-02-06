package services

import (
	"bookstore-api/src/models"
	_ "bookstore-api/src/models"
	"errors"
	_ "errors"
)

var books = []models.Book{}

func AddBook(book models.Book) {
	books = append(books, book)
}

func GetBooks() []models.Book {
	return books
}

func GetBookByID(id string) (models.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}

func DeleteBook(id string) error {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
