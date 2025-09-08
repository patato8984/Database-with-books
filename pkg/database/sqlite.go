package database

import (
	"database/sql"

	"os"

	_ "modernc.org/sqlite"
)

func NewSQLiteConnection(db_patch string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", db_patch)
	if err != nil {
		return nil, err
	}
	bytes, er := os.ReadFile("create_table.sql")
	if er != nil {
		return nil, er
	}
	base := string(bytes)
	_, errr := db.Exec(base)
	if errr != nil {
		return nil, errr
	}
	return db, nil
}
