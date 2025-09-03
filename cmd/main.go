package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

type Books struct {
	Author    string `json:"author"`
	Id_author int    `json:"id_author"`
	Id        int
	Name      string `json:"name"`
	Age       int    `json:"age"`
}
type Books_Get struct {
	Id_author  int
	Id_books   int
	Name_books string
	Age        int
}
type NewAuthor struct {
	Name string `json:"name"`
}

func Delbooks(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "ID должен быть числом"}`, http.StatusBadRequest)
		return
	}
	_, er := db.Query("DELETE FROM books WHERE ID = ?", id)
	if er != nil {
		http.Error(w, `{"status":"not found id"}`, http.StatusBadRequest)
		log.Print(er)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{"status" : "delete"}`)

}
func books(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	switch r.Method {
	case http.MethodGet:
		rows, err := db.Query("SELECT id_books, id_author, name_books, age FROM books")
		if err != nil {
			http.Error(w, `{"status":"error in server"}`, http.StatusInternalServerError)
		}
		defer rows.Close()
		var bok []Books_Get
		for rows.Next() {
			var b Books_Get
			if err := rows.Scan(&b.Id_books, &b.Id_author, &b.Name_books, &b.Age); err != nil {
				http.Error(w, `{"status":"error"}`, http.StatusBadRequest)
				log.Print(err)
				return
			}
			bok = append(bok, b)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bok)
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		var newBooks Books
		if err := json.NewDecoder(r.Body).Decode(&newBooks); err != nil {
			http.Error(w, `{"status":"error"}`, http.StatusBadRequest)
			log.Print(err)
			return
		}
		_, err := db.Exec("INSERT INTO books (id_author, name_books, age) VALUES (?, ?, ?)", newBooks.Id_author, newBooks.Name, newBooks.Age)
		if err != nil {
			http.Error(w, `{"status":"error"}`, http.StatusNotFound)
			log.Print(err)
			return
		}
		json.NewEncoder(w).Encode(`{"status":"Updated"}`)
	}
}
func author(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		var a NewAuthor
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			http.Error(w, `{"status": "error"}`, http.StatusBadRequest)
			return
		}
		_, err := db.Exec("INSERT INTO author (name) VALUES (?)", a.Name)
		if err != nil {
			http.Error(w, `{"status": "error"}`, http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(`{"status":"Updated"}`)
	}
}
func Delauthor(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	switch r.Method {
	case http.MethodDelete:
		idStr := strings.TrimPrefix(r.URL.Path, "/books/author/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, `{"error": "ID должен быть числом"}`, http.StatusBadRequest)
			return
		}
		_, er := db.Query("DELETE FROM books WHERE id_author = ?", id)
		if er != nil {
			http.Error(w, `{"status":"error"}`, http.StatusBadRequest)
		}
		json.NewEncoder(w).Encode(`{"status":"delete"}`)
	}
}
func main() {
	db, err := sql.Open("sqlite", "books.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	bytes, err := os.ReadFile("create_tables.sql")
	if err != nil {
		log.Panic(err)
	}
	base := string(bytes)
	_, er := db.Exec(base)
	if er != nil {
		log.Panic(er)
	}
	log.Print("База данных создана")
	http.HandleFunc("/books/author/", Delauthor)
	http.HandleFunc("/books/author", author)
	http.HandleFunc("/books/", Delbooks)
	http.HandleFunc("/books", books)
	log.Print("Сервер запущен")
	http.ListenAndServe(":8080", nil)
}
