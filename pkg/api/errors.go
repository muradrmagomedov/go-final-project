package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	errorMarshal       = "Ошибка сериализации данных"
	errorUnmarshal     = "Ошибка десериализации данных"
	errorEmpty         = "Не задан параметр: "
	errorWrongFormat   = "Некорректно задан параметр: "
	errorReadData      = "Ошибка чтения данных"
	errorNotFound      = "Данные не найдены"
	errorWrongPassword = "Неверный пароль"
	errorNotAuthorized = "Пользователь не авторизован"
)

type Error struct {
	Message string `json:"error"`
}

func errorMessage(path, message string) *Error {
	return &Error{Message: fmt.Sprintf("%s:%s", path, message)}
}

func sendError(w http.ResponseWriter, op string, message string, status int) {
	errorMessage := errorMessage(op, message)
	logrus.Error(errorMessage)
	jsonMsg, err := json.Marshal(errorMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(jsonMsg)
}
