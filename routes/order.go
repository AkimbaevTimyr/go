package routes

import (
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func OrderMux() http.Handler {
	mux := http.NewServeMux()

	c := injection.InitOrderController()

	mux.Handle("/create", middleware.CreateMiddleware(middleware.AuthMiddleware, middleware.CheckAdmin)(http.HandlerFunc(c.CreateOrder)))
	mux.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(c.GetOrders)))
	mux.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(c.Delete)))
	mux.Handle("/change-status", middleware.CreateMiddleware(middleware.AuthMiddleware, middleware.CheckModerator)(http.HandlerFunc(c.ChangeStatus)))

	return mux
}
