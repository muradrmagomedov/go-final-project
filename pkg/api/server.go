package server

import (
	"fmt"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/muradrmagomedov/final-project/pkg/repository"
	"github.com/sirupsen/logrus"
)

var Repo *repository.Repository = repository.NewRepository()

type Server struct {
}

func (s *Server) Run(port string) error {
	const op = "server.Run"
	logrus.Info("Starting server on port:", port)
	router := chi.NewRouter()
	fs := http.FileServer(http.Dir("web"))
	router.Handle("/*", http.StripPrefix("/", fs))

	router.Get("/api/nextdate", nextDate)
	router.Get("/api/task", auth(getTask))
	router.Post("/api/task", auth(addTask))
	router.Put("/api/task", auth(updateTask))
	router.Delete("/api/task", auth(deleteTask))
	router.Post("/api/task/done", auth(repeatTask))
	router.Get("/api/tasks", auth(allTasks))
	router.Post("/api/signin", signIn)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return fmt.Errorf("Ошибка запуска сервера. %s:%v", op, err)
	}
	return nil
}
