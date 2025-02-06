package main

import (
	"bookstore-api/src/controllers"
	"fmt"
	_ "fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetBooks(w)
		case http.MethodPost:
			controllers.CreateBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/book", controllers.GetBook)
	http.HandleFunc("/book/delete", controllers.DeleteBook)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
