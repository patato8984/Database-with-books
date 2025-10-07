package auth

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (b *AuthRepository) CreateUser(login, password string) error {
	_, err := b.db.Exec("INSERT INTO users (user_name, user_password) VALUES ($1, $2)", login, password)
	if err != nil {
		return err
	}
	return nil
}
func (b *AuthRepository) GetHashPassworld(login string) (string, error) {
	var l string
	rows, err := b.db.Query("SELECT user_password FROM users WHERE user_name = $1", login)
	if err != nil {
		return l, errors.New("the user was not found")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&l)
		if err != nil {
			return l, err
		}
	}
	return l, nil
}
