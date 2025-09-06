package books

import (
	"encoding/json"
	"log"
	"myapp/internal/models"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	bookService *BookService
}

func NewBookHandler(bookService *BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}
func (s *BookHandler) DelAuthor(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/author/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"status":"error url"}`, http.StatusBadRequest)
		return
	}
	if er := s.bookService.DelAuthor(id); er != nil {
		http.Error(w, `{"status":"error db"}`, http.StatusInternalServerError)
		log.Print(er)
		return
	}
}
func (s *BookHandler) DelBooks(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"status":"error url"}`, http.StatusBadRequest)
		return
	}
	if er := s.bookService.DelBook(id); er != nil {
		http.Error(w, `{"status":"error db"}`, http.StatusInternalServerError)
		log.Print(er)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{"status":"deleted!"}`)
}

func (s *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := s.bookService.GetAllBooks()
	if err != nil {
		http.Error(w, `{"status":"error db"}`, http.StatusInternalServerError)
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func (s *BookHandler) GetAllBooksAuthor(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/author/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"status":"error url"}`, http.StatusBadRequest)
		return
	}
	books, err := s.bookService.GetAllBooksAuthor(id)
	if err != nil {
		http.Error(w, `{"status":"error db"}`, http.StatusInternalServerError)
		log.Print(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func (s *BookHandler) CreateAuthorBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Books
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, `{"status":"error json"}`, http.StatusBadRequest)
		log.Print(err)
		return
	}
	er := s.bookService.CreateAuthor(book)
	if er != nil {
		http.Error(w, `{"status":"error db"}`, http.StatusInternalServerError)
		log.Print(er)
		return
	}
}
func (s *BookHandler) CreateBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Books
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, `{"status":"error json"}`, http.StatusBadRequest)
		return
	}
	if er := s.bookService.CreateBooks(book); er != nil {
		http.Error(w, `{"status":"error"}`, http.StatusInternalServerError)
		log.Print(er)
		return
	}
}
