package server

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/muradrmagomedov/final-project/pkg/services"
	"github.com/muradrmagomedov/final-project/pkg/todo"
)

func updateTask(w http.ResponseWriter, r *http.Request) {

	const op = "server.updateTask"
	var newTask todo.Task

	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(w, op, errorReadData, http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &newTask)
	if err != nil {
		sendError(w, op, errorUnmarshal, http.StatusInternalServerError)
		return
	}
	if newTask.Title == "" {
		msg := errorEmpty + "title"
		sendError(w, op, msg, http.StatusBadRequest)
		return
	}

	if newTask.Date == "" {
		newTask.Date = time.Now().Format("20060102")
	}

	date, err := time.Parse("20060102", newTask.Date)
	if err != nil {
		msg := errorWrongFormat + "date"
		sendError(w, op, msg, http.StatusBadRequest)
		return
	}
	if date.Unix() < time.Now().Unix() {
		if newTask.Repeat == "" {
			date = time.Now()
			newTask.Date = date.Format("20060102")
		} else {
			nextDate, err := services.NextDate(time.Now(), newTask.Date, newTask.Repeat)
			if err != nil {
				sendError(w, op, err.Error(), http.StatusBadRequest)
				return
			}
			newTask.Date = nextDate
		}
	}
	err = Repo.UpdateTask(newTask)
	if err != nil {
		sendError(w, op, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	answerId, err := json.Marshal(struct{}{})
	if err != nil {
		sendError(w, op, errorMarshal, http.StatusInternalServerError)
		return
	}
	w.Write(answerId)
}
