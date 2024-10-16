package main

import (
	"akimbaev/database"
	"akimbaev/routes"
	"log"
	"net/http"
)

func main() {
	database.Init()
	database.RedisInit()
	api := http.NewServeMux()
	api.Handle("/api/v1/user/", http.StripPrefix("/api/v1/user", routes.UserMux()))
	api.Handle("/api/v1/auth/", http.StripPrefix("/api/v1/auth", routes.AuthMux()))
	api.Handle("/api/v1/post/", http.StripPrefix("/api/v1/post", routes.OrderMux()))
	api.Handle("/api/v1/category/", http.StripPrefix("/api/v1/category", routes.CategoryMux()))
	api.Handle("/api/v1/subscription/", http.StripPrefix("/api/v1/subscription", routes.SubscriptionMux()))
	log.Println("Listening on :8080")

	http.ListenAndServe(":8080", api)
}
