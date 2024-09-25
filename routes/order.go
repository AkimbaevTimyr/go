package routes

import (
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func OrderMux() http.Handler {
	orderMux := http.NewServeMux()

	controller := injection.InitOrderController()

	orderMux.Handle("/create", middleware.CreateMiddleware(middleware.AuthMiddleware, middleware.CheckAdmin)(http.HandlerFunc(controller.CreateOrder)))
	orderMux.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(controller.GetOrders)))
	orderMux.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(controller.Delete)))
	orderMux.Handle("/change-status", middleware.CreateMiddleware(middleware.AuthMiddleware, middleware.CheckModerator)(http.HandlerFunc(controller.ChangeStatus)))

	return orderMux
}
