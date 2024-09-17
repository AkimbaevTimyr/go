package main

import (
	"akimbaev/database"
	"akimbaev/handlers"
	"akimbaev/response"
	"akimbaev/routes"
	"log"
	"net/http"
)

var secretKey = []byte("secret-key")

func main() {
	database.Init()
	api := http.NewServeMux()
	api.Handle("/api/v1/user/", http.StripPrefix("/api/v1/user", routes.UserMux()))
	log.Println("Listening on :8080")

	http.ListenAndServe(":8080", api)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	response.Json(w, http.StatusUnauthorized, "Invalid token")

	tokenString = tokenString[len("Bearer "):]

	err := handlers.VerifyToken(tokenString)

	if err != nil {
		response.Json(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	return
}
