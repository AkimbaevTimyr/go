package routes

import (
	"akimbaev/handlers"
	"akimbaev/middleware"
	"net/http"
)

func OrderMux() http.Handler {
	orderMux := http.NewServeMux()

	orderMux.Handle("/create", middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateOrder)))
	orderMux.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetOrders)))

	return orderMux
}
