package moks

import "github.com/stretchr/testify/mock"

type MockAuthRepository struct {
	mock.Mock
}

func (m *MockAuthRepository) CreateUser(login, password string) error {
	args := m.Called(login, password)
	return args.Error(0)
}
func (m *MockAuthRepository) GetHashPassworld(login string) (string, error) {
	args := m.Called(login)
	return "", args.Error(0)
}
