package main

import (
	"log"
	"myapp/internal/books"
	"myapp/pkg/database"
	"net/http"
)

func main() {
	db, err := database.NewSQLiteConnection()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	newRepository := books.NewBookRepository(db)
	newServise := books.NewBookService(newRepository)
	newHandler := books.NewBookHandler(newServise)
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			newHandler.GetAllBooks(w, r)
		case http.MethodPost:
			newHandler.CreateBooks(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			newHandler.DelBooks(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})
	http.HandleFunc("/books/author", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			newHandler.CreateAuthorBooks(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})
	http.HandleFunc("/books/author/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			newHandler.GetAllBooksAuthor(w, r)
		case http.MethodDelete:
			newHandler.DelAuthor(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})
	log.Print("Яица")
	http.ListenAndServe(":8080", nil)
}
