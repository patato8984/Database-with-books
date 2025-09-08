package main

import (
	"log"
	"myapp/internal/auth"
	"myapp/internal/books"
	"myapp/pkg/config"
	"myapp/pkg/database"
	"myapp/pkg/middleware"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Panic(err)
	}
	db, er := database.NewSQLiteConnection(cfg.DBpatch)
	if er != nil {
		log.Panic(err)
	}
	defer db.Close()
	newRepository := books.NewBookRepository(db)
	newServise := books.NewBookService(newRepository)
	newHandler := books.NewBookHandler(newServise)

	newRepositoryAuth := auth.NewAuthRepository(db)
	newServiseAuth := auth.NewAuthService(newRepositoryAuth, cfg.Jwt)
	newHandlerAuth := auth.NewAuthHandler(newServiseAuth)

	newMiddlewareAuth := middleware.NewAuthorization(cfg.Jwt)
	http.HandleFunc("/authentication", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			newHandlerAuth.Authentication(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			newHandlerAuth.NewRegister(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			newHandler.GetAllBooks(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})
	http.Handle("/books/create", newMiddlewareAuth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			newHandler.CreateBooks(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})))
	http.Handle("/books/", newMiddlewareAuth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			newHandler.DelBooks(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})))
	http.Handle("/books/author", newMiddlewareAuth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			newHandler.CreateAuthorBooks(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})))
	http.HandleFunc("/books/author/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			newHandler.GetAllBooksAuthor(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})
	http.Handle("/books/author/delete/", newMiddlewareAuth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodDelete:
			newHandler.DelAuthor(w, r)
		default:
			http.Error(w, `{"status":"error method"}`, http.StatusBadRequest)
		}
	})))
	log.Print("Яица")
	http.ListenAndServe(":8080", nil)
}
