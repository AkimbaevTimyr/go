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
	orderMux.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteOrder)))
	orderMux.Handle("/change-status", middleware.AuthMiddleware(http.HandlerFunc(handlers.ChangeStatus)))

	return orderMux
}
