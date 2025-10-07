package books

import (
	"database/sql"
	"myapp/internal/models"

	_ "github.com/lib/pq"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}
func (b *BookRepository) DelAuthors(id int) error {
	_, err := b.db.Exec("DELETE FROM books WHERE id_author = ?", id)
	if err != nil {
		return err
	}
	return nil
}
func (b *BookRepository) DelBooks(id int) error {
	_, err := b.db.Exec("DELETE FROM books WHERE id_books = ?", id)
	if err != nil {
		return err
	}
	return nil
}
func (b *BookRepository) GetAllBooks() ([]models.Books, error) {
	var books []models.Books
	rows, err := b.db.Query("SELECT b.id_books, b.name_books, b.age, b.id_author AS id_author, a.name AS name_author FROM books b JOIN author a ON b.id_author = a.id")
	if err != nil {
		return books, err
	}
	defer rows.Close()
	for rows.Next() {
		var book models.Books
		if er := rows.Scan(&book.Id, &book.Name, &book.Age, &book.Id_author, &book.Name_author); er != nil {
			return books, er
		}
		books = append(books, book)
	}
	return books, nil
}
func (b *BookRepository) GetAllBooksAuthor(id int) ([]models.Books, error) {
	var books []models.Books
	rows, err := b.db.Query("SELECT a.id, a.name, b.id_books AS id_books, b.name_books AS name_books, b.age AS age FROM author a JOIN books b ON a.id = b.id_author WHERE a.id = ?", id)
	if err != nil {
		return books, err
	}
	defer rows.Close()
	for rows.Next() {
		var book models.Books
		if err := rows.Scan(&book.Id_author, &book.Name_author, &book.Id, &book.Name, &book.Age); err != nil {
			return books, err
		}
		books = append(books, book)
	}
	return books, nil
}
func (b *BookRepository) CreateAuthor(bo models.Books) error {
	_, err := b.db.Exec("INSERT INTO author (name) VALUES ($1)", bo.Name_author)
	if err != nil {
		return err
	}
	return nil
}
func (b *BookRepository) CreateBooks(bo models.Books) error {
	_, err := b.db.Exec("INSERT INTO books (id_author, name_books, age) VALUES ($1, $2, $3)", bo.Id_author, bo.Name, bo.Age)
	if err != nil {
		return err
	}
	return nil
}
