package server

import (
	"encoding/json"
	"net/http"
)

func deleteTask(w http.ResponseWriter, r *http.Request) {

	const op = "server.deleteTask"

	id := r.URL.Query().Get("id")
	if id == "" {
		msg := errorEmpty + "id"
		sendError(w, op, msg, http.StatusBadRequest)
		return
	}
	err := Repo.DeleteTask(id)
	if err != nil {
		sendError(w, op, err.Error(), http.StatusInternalServerError)
		return
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
