package routes

import (
	"akimbaev/handlers"
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func UserMux() http.Handler {
	userMux := http.NewServeMux()

	userController := injection.InitUserController()

	userMux.Handle("/get", middleware.AuthMiddleware(http.HandlerFunc(userController.GetUser)))
	userMux.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(userController.DeleteUser)))
	userMux.Handle("/update", middleware.AuthMiddleware(http.HandlerFunc(userController.UpdateUser)))

	userMux.Handle("/login", http.HandlerFunc(handlers.Login))
	userMux.Handle("/register", http.HandlerFunc(handlers.Register))
	userMux.Handle("/checkCode", http.HandlerFunc(handlers.CheckCode))
	return userMux
}
