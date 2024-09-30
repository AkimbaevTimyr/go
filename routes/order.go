package routes

import (
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func OrderMux() http.Handler {
	mux := http.NewServeMux()

	c := injection.InitOrderController()

	mux.Handle("/create", middleware.CreateMiddleware(middleware.AuthMiddleware, middleware.CheckAdmin)(http.HandlerFunc(c.CreatePost)))
	mux.Handle("/orders", middleware.AuthMiddleware(http.HandlerFunc(c.GetPosts)))
	mux.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(c.Delete)))

	return mux
}
