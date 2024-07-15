package server

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/muradrmagomedov/final-project/pkg/services"
)

type jwtToken struct {
	Token string `json:"token"`
}

type user struct {
	Password string `json:"password"`
}

func signIn(w http.ResponseWriter, r *http.Request) {
	const op = "server.signIn"
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendError(w, op, errorReadData, http.StatusInternalServerError)
		return
	}
	userPassword := os.Getenv("TODO_PASSWORD")
	var user user
	err = json.Unmarshal(body, &user)
	if err != nil {
		sendError(w, op, errorUnmarshal, http.StatusInternalServerError)
		return
	}
	if userPassword != user.Password {
		sendError(w, op, errorWrongPassword, http.StatusBadRequest)
		return
	}
	token, err := services.Auth(user.Password)
	if err != nil {
		sendError(w, op, err.Error(), http.StatusInternalServerError)
		return
	}
	jwtToken := jwtToken{Token: token}
	ans, _ := json.Marshal(jwtToken)
	w.Write(ans)
}
