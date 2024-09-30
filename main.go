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
	api.Handle("/api/v1/order/", http.StripPrefix("/api/v1/order", routes.OrderMux()))
	api.Handle("/api/v1/report/", http.StripPrefix("/api/v1/report", routes.OrderReportMux()))
	api.Handle("/api/v1/subscription/", http.StripPrefix("/api/v1/subscription", routes.SubscriptionMux()))
	log.Println("Listening on :8080")

	http.ListenAndServe(":8080", api)
}
