package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

const (
	schedulerTable = "scheduler"
)

// CreateSchedulerTable() creates schedule table in database if it is not existed
func (r *Repository) CreateSchedulerTable(db *sqlx.DB) error {
	const op = "repository.CreateSchedulerTable"

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
	id INTEGER NOT NULL PRIMARY KEY,
	date VARCHAR(8) NOT NULL,
	title VARCHAR(255) NOT NULL,
	comment VARCHAR(255) NOT NULL,
	repeat VARCHAR(128) NOT NULL
	);`, schedulerTable)

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("Запрос на создание таблицы не выполнен. %s:%v", op, err)
	}
	return nil
}

// CreateIndex() creates index on attribute (date) of table schedule
func (r *Repository) CreateIndex(db *sqlx.DB) error {
	const op = "repository.CreateIndex"
	query := fmt.Sprintf("CREATE INDEX IF NOT EXISTS date_index ON %s (date);", schedulerTable)
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("Запрос на создание индекса не выполнен. %s:%v", op, err)
	}
	return nil
}

// CreateDB() create database file using provided path if it is not created
func (r *Repository) CreateDBFile(dbPath string) error {
	const op = "repository.CreateDBFile"
	if _, err := os.Stat(dbPath); err != nil {
		logrus.Info("Файл с базой данных не найден, создаем новый файл.")
		_, err = os.Create(dbPath)
		if err != nil {
			return fmt.Errorf("Не удалось создать файл с базой данных: %s:%v", op, err)
		}
	}
	return nil
}

// InitiateDB() creates or opens database, creating necessary tables and indexes
func (r *Repository) InitiateDB(dbPath string) error {
	err := r.CreateDBFile(dbPath)
	if err != nil {
		return err
	}
	err = r.SqliteDB(dbPath)
	if err != nil {
		return err
	}
	err = r.CreateSchedulerTable(r.db)
	if err != nil {
		return err
	}
	err = r.CreateIndex(r.db)
	if err != nil {
		return err
	}
	logrus.Info("База данных была успешно создана и инициализирована.")
	return nil

}
