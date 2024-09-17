package middleware

import (
	"akimbaev/handlers"
	"akimbaev/response"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		tokenString = tokenString[len("Bearer "):]

		if tokenString == "" {
			response.Json(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		err := handlers.VerifyToken(tokenString)

		if err != nil {
			response.Json(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
