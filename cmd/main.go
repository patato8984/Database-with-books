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
	Id   int
	Name string `json:"name"`
	Age  int    `json:"age"`
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
		rows, err := db.Query("SELECT id, name, age FROM books")
		if err != nil {
			http.Error(w, `{"status":"error in server"}`, http.StatusInternalServerError)
		}
		defer rows.Close()
		var bok []Books
		for rows.Next() {
			var b Books
			if err := rows.Scan(&b.Id, &b.Name, &b.Age); err != nil {
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
		}
		_, err := db.Exec("INSERT INTO books (name, age) VALUES (?, ?)", newBooks.Name, newBooks.Age)
		if err != nil {
			http.Error(w, `{"status":"error"}`, http.StatusNotFound)
		}
		json.NewEncoder(w).Encode(`{"status":"Updated"}`)
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
	http.HandleFunc("/books/", Delbooks)
	http.HandleFunc("/books", books)
	log.Print("Сервер запущен")
	http.ListenAndServe(":8080", nil)
}
