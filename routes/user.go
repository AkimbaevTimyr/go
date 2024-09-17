package routes

import (
	"akimbaev/handlers"
	"akimbaev/middleware"
	"net/http"
)

func UserMux() http.Handler {
	userMux := http.NewServeMux()
	userMux.Handle("/get", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetUser)))
	userMux.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteUser)))
	userMux.Handle("/update", middleware.AuthMiddleware(http.HandlerFunc(handlers.UpdateUser)))

	userMux.Handle("/login", http.HandlerFunc(handlers.Login))
	userMux.Handle("/register", http.HandlerFunc(handlers.Register))
	userMux.Handle("/checkCode", http.HandlerFunc(handlers.CheckCode))
	return userMux
}
