package middleware

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/response"
	"fmt"
	"net/http"
)

func CheckSubscription(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userClaims, _ := helpers.ExctractUserFromToken(r)

		redisInstance := database.Redis{}
		val, err := redisInstance.Get(fmt.Sprintf("%v", userClaims.UserID))

		if err != nil {
			response.Json(w, http.StatusForbidden, helpers.Envelope{"message": "Subscription not found"})
			return
		}

		if val == "0" {
			response.Json(w, http.StatusForbidden, helpers.Envelope{"message": "Subscription not found"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
