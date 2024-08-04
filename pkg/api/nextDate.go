package server

import (
	"net/http"
	"time"

	"github.com/muradrmagomedov/final-project/pkg/services"
)

func nextDate(w http.ResponseWriter, r *http.Request) {
	const op = "server.nextDate"

	now := r.URL.Query().Get("now")
	date := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	timeNow, err := time.Parse("20060102", now)
	if err != nil {
		msg := errorWrongFormat + "now"
		sendError(w, op, msg, http.StatusBadRequest)
	}
	newDate, err := services.NextDate(timeNow, date, repeat)
	if err != nil {
		sendError(w, op, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(newDate))
}
