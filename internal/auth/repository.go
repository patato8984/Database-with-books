package auth

import (
	"database/sql"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (b *AuthRepository) CreateUser(login, password string) error {
	_, err := b.db.Exec("INSERT INTO user (user_name, user_password) VALUES (?, ?)", login, password)
	if err != nil {
		return err
	}
	return nil
}
func (b *AuthRepository) GetHashPassworld(login string) (string, error) {
	var l string
	rows, err := b.db.Query("SELECT user_password FROM user WHERE user_name = ?", login)
	if err != nil {
		return l, err
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
