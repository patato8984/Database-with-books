package books

import (
	"errors"
	"myapp/internal/models"
)

type BookService struct {
	bookRapo *BookRepository
}

func NewBookService(bookRapo *BookRepository) *BookService {
	return &BookService{bookRapo: bookRapo}
}
func (r *BookService) DelAuthor(id int) error {
	if id <= 0 {
		return errors.New("the id must not be equal to or less than zero")
	}
	if err := r.bookRapo.DelAuthor(id); err != nil {
		return err
	}
	return nil
}
func (r *BookService) DelBook(id int) error {
	if id <= 0 {
		return errors.New("the id must not be equal to or less than zero")
	}
	if err := r.bookRapo.DelBooks(id); err != nil {
		return err
	}
	return nil
}
func (r *BookService) GetAllBooks() ([]models.Books, error) {
	var books []models.Books
	GetAllBooks, err := r.bookRapo.GetAllBooks(books)
	if err != nil {
		return GetAllBooks, err
	}
	return GetAllBooks, nil
}
func (r *BookService) GetAllBooksAuthor(id int) ([]models.Books, error) {
	var books []models.Books
	GetALLBooks, err := r.bookRapo.GetAllBooksAuthor(books, id)
	if err != nil {
		return books, err
	}
	return GetALLBooks, nil
}
func (r *BookService) CreateAuthor(book models.Books) error {
	if book.Name_author == "" {
		return errors.New("dsa")
	}
	err := r.bookRapo.CreateAuthor(book)
	if err != nil {
		return err
	}
	return nil
}
func (r *BookService) CreateBooks(book models.Books) error {
	if book.Name == "" {
		return errors.New("dsa")
	} else if book.Age <= 0 {
		return errors.New("dsa")
	}
	if err := r.bookRapo.CreateBooks(book); err != nil {
		return err
	}
	return nil
}
