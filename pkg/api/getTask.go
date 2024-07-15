package server

import (
	"encoding/json"
	"net/http"
)

func getTask(w http.ResponseWriter, r *http.Request) {

	const op = "server.getTask"

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

	jsnTask, err := json.Marshal(task)
	if err != nil {
		sendError(w, op, errorMarshal, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsnTask)

}
