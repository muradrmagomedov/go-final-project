package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/muradrmagomedov/final-project/pkg/services"
)

func repeatTask(w http.ResponseWriter, r *http.Request) {

	const op = "server.repeatTask"

	id := r.URL.Query().Get("id")
	if id == "" {
		msg := errorEmpty + "id"
		sendError(w, op, msg, http.StatusBadRequest)
		return
	}

	task, err := Repo.GetTaskById(id)
	if err != nil {
		sendError(w, op, err.Error(), http.StatusBadRequest)
		return
	}

	if task.Repeat == "" {
		err = Repo.DeleteTask(id)
		if err != nil {
			sendError(w, op, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		date, err := services.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			sendError(w, op, err.Error(), http.StatusInternalServerError)
			return
		}
		task.Date = date
		err = Repo.UpdateTask(*task)
		if err != nil {
			sendError(w, op, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	answer, err := json.Marshal(struct{}{})
	if err != nil {
		sendError(w, op, errorMarshal, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(answer)

}
