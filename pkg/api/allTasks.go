package server

import (
	"encoding/json"
	"net/http"

	"github.com/muradrmagomedov/final-project/pkg/todo"
)

func allTasks(w http.ResponseWriter, r *http.Request) {

	const op = "handlers.tasks"
	var tasks todo.Tasks

	query := r.URL.Query().Get("search")
	tasks, err := Repo.GetAllTasks(query)
	if err != nil {
		sendError(w, op, err.Error(), http.StatusInternalServerError)
		return
	}

	jsnTasks, err := json.Marshal(tasks)
	if err != nil {
		sendError(w, op, errorMarshal, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsnTasks)
}
