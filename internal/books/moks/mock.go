package moks

import (
	"myapp/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) CreateAuthor(book models.Books) error {
	args := m.Called(book)
	return args.Error(0)
}
func (m *MockBookRepository) CreateBooks(book models.Books) error {
	args := m.Called(book)
	return args.Error(0)
}
func (m *MockBookRepository) DelAuthors(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockBookRepository) DelBooks(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockBookRepository) GetAllBooks() ([]models.Books, error) {
	i := []models.Books{}
	args := m.Called()
	return i, args.Error(0)
}
func (m *MockBookRepository) GetAllBooksAuthor(id int) ([]models.Books, error) {
	i := []models.Books{}
	args := m.Called()
	return i, args.Error(0)
}
