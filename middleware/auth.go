package middleware

import (
	"akimbaev/helpers"
	"akimbaev/response"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		//TODO если токен не передан ошибка
		tokenString = tokenString[len("Bearer "):]

		if tokenString == "" {
			response.Json(w, http.StatusUnauthorized, map[string]string{
				"message": "Invalid token",
			})
			return
		}

		if err := helpers.VerifyToken(tokenString); err != nil {
			response.Json(w, http.StatusUnauthorized, map[string]string{
				"message": "Invalid token",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
