package books

import (
	"errors"
	"myapp/internal/models"
)

type bookRepository interface {
	CreateAuthor(book models.Books) error
	CreateBooks(book models.Books) error
	DelAuthors(id int) error
	DelBooks(id int) error
	GetAllBooks() ([]models.Books, error)
	GetAllBooksAuthor(id int) ([]models.Books, error)
}

type BookService struct {
	bookRapo bookRepository
}

func NewBookService(BookRapo bookRepository) *BookService {
	return &BookService{bookRapo: BookRapo}
}
func (r *BookService) DelAuthor(id int) error {
	if id <= 0 {
		return errors.New("the id must not be equal to or less than zero")
	}
	if err := r.bookRapo.DelAuthors(id); err != nil {
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
	GetAllBooks, err := r.bookRapo.GetAllBooks()
	if err != nil {
		return GetAllBooks, err
	}
	return GetAllBooks, nil
}
func (r *BookService) GetAllBooksAuthor(id int) ([]models.Books, error) {

	GetALLBooks, err := r.bookRapo.GetAllBooksAuthor(id)
	if err != nil {
		return GetALLBooks, err
	}
	return GetALLBooks, nil
}
func (r *BookService) CreateAuthor(book models.Books) error {
	if book.Name_author == "" {
		return errors.New("error json")
	}
	err := r.bookRapo.CreateAuthor(book)
	if err != nil {
		return err
	}
	return nil
}
func (r *BookService) CreateBooks(book models.Books) error {
	if book.Name == "" {
		return errors.New("error json")
	} else if book.Age <= 0 {
		return errors.New("error json")
	}
	if err := r.bookRapo.CreateBooks(book); err != nil {
		return err
	}
	return nil
}
