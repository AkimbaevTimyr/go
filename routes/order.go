package routes

import (
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func OrderMux() http.Handler {
	orderMux := http.NewServeMux()

	controller := injection.InitOrderController()

	orderMux.Handle("/create", middleware.AuthMiddleware(http.HandlerFunc(controller.CreateOrder)))
	orderMux.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(controller.GetOrders)))
	// orderMux.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteOrder)))
	// orderMux.Handle("/change-status", middleware.AuthMiddleware(http.HandlerFunc(handlers.ChangeStatus)))

	return orderMux
}
