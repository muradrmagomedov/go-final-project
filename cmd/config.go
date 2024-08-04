package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	dbPath string
	port   string
}

func loadConfig() (*Config, error) {
	const op = "main.loadConfig"
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Не удалось загрузить переменные окружения из файла .env. %s:%v", op, err)
	}
	var config Config
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
		logrus.Warnf("Проблема с загрузкой порта из .env. Загружен порт по-умолчанию:%s", port)
	}
	db_path := os.Getenv("DB_PATH")
	config.port = port
	config.dbPath = db_path
	return &config, nil
}
