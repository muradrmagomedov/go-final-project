package server

import (
	"net/http"
	"os"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "server.auth"
		// смотрим наличие пароля
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var jwt string // JWT-токен из куки
			// получаем куку
			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}
			var valid bool
			// здесь код для валидации и проверки JWT-токена
			authToken := os.Getenv("JWT_TOKEN")
			valid = authToken == jwt
			if !valid {
				// возвращаем ошибку авторизации 401
				sendError(w, op, errorNotAuthorized, http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}
