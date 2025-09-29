package books_test

import (
	"errors"
	"myapp/internal/books"
	"myapp/internal/books/moks"
	"myapp/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServiseCreateBooks(t *testing.T) {
	tests := []struct {
		book       models.Books
		shauldCall bool
		errors     error
	}{
		{
			book: models.Books{
				Name:      "dasd",
				Age:       1999,
				Id_author: 1,
			},
			shauldCall: true,
			errors:     nil,
		},
		{
			book: models.Books{
				Name:      "",
				Age:       1999,
				Id_author: 1,
			},
			shauldCall: false,
			errors:     errors.New("error json"),
		},
		{
			book: models.Books{
				Name:      "dsa",
				Age:       0,
				Id_author: 1,
			},
			shauldCall: false,
			errors:     errors.New("error json"),
		},
	}
	for _, tt := range tests {
		t.Run("popka", func(t *testing.T) {
			mockRapo := new(moks.MockBookRepository)
			servis := books.NewBookService(mockRapo)
			if tt.shauldCall {
				mockRapo.On("CreateBooks", tt.book).Return(nil)
			}
			err := servis.CreateBooks(tt.book)
			if tt.errors == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, err.Error(), tt.errors.Error())
				assert.Error(t, err)
			}
			if tt.shauldCall {
				mockRapo.AssertCalled(t, "CreateBooks", tt.book)
			} else {
				mockRapo.AssertNotCalled(t, "CreateBooks")
			}
		})
	}
}
func Test_ServiseDelAuthor(t *testing.T) {
	tests := []struct {
		id         int
		errors     error
		shouldCall bool
	}{
		{
			id:         1,
			errors:     nil,
			shouldCall: true,
		},
		{
			id:         24,
			errors:     nil,
			shouldCall: true,
		},
		{
			id:         0,
			errors:     errors.New("the id must not be equal to or less than zero"),
			shouldCall: false,
		},
		{
			id:         -432,
			errors:     errors.New("the id must not be equal to or less than zero"),
			shouldCall: false,
		},
	}
	for _, tt := range tests {
		t.Run("kaki", func(t *testing.T) {
			mockRapo := new(moks.MockBookRepository)
			testServise := books.NewBookService(mockRapo)
			if tt.shouldCall {
				mockRapo.On("DelAuthors", tt.id).Return(tt.errors)
			}
			err := testServise.DelAuthor(tt.id)
			if tt.errors == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.errors.Error(), err.Error())
			}
			if tt.shouldCall {
				mockRapo.AssertCalled(t, "DelAuthors", tt.id)
			} else {
				mockRapo.AssertNotCalled(t, "DelAuthors")
			}
		})
	}
}
