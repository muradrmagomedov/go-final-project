package main

import (
	"github.com/mattn/go-colorable"
	server "github.com/muradrmagomedov/final-project/pkg/api"
	"github.com/muradrmagomedov/final-project/pkg/repository"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())

	config, err := loadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.NewRepository()
	err = repo.InitiateDB(config.dbPath)
	if err != nil {
		logrus.Fatal(err)
	}

	server.Repo = repo
	server.Repo.SqliteDB(config.dbPath)

	srv := new(server.Server)

	err = srv.Run(config.port)
	if err != nil {
		logrus.Fatal("Cant start server:", err)
	}
}
