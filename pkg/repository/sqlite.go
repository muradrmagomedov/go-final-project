package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type Repository struct {
	db *sqlx.DB
}

// NewRepository() creates new object to control work with database
func NewRepository() *Repository {
	return &Repository{}
}

// SqliteDB() connect to sqlite database using provided path
func (r *Repository) SqliteDB(dbPath string) error {
	const op = "repository.SqliteDB"
	db, err := sqlx.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("Соединение с базой данных не было установлено. %s:%v", op, err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("Соединение с базой данных недоступно. %s:%v", op, err)
	}
	r.db = db
	return nil
}
