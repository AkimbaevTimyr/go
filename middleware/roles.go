package middleware

import (
	"akimbaev/helpers"
	"akimbaev/response"
	"net/http"
)

func CheckAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userClaims, _ := helpers.ExctractUserFromToken(r)

		if userClaims.Role != "admin" {
			response.Json(w, http.StatusForbidden, map[string]string{
				"message": "Forbidden",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func CheckModerator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userClaims, _ := helpers.ExctractUserFromToken(r)

		if userClaims.Role != "moderator" && userClaims.Role != "admin" {
			response.Json(w, http.StatusForbidden, map[string]string{
				"message": "Forbidden",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
