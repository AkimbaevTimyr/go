package routes

import (
	"akimbaev/controller"
	"akimbaev/handlers"
	"akimbaev/middleware"
	"akimbaev/repository"
	"akimbaev/service"
	"net/http"
)

func UserMux() http.Handler {
	userMux := http.NewServeMux()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	userMux.Handle("/get", middleware.AuthMiddleware(http.HandlerFunc(userController.GetUser)))
	userMux.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(userController.DeleteUser)))
	userMux.Handle("/update", middleware.AuthMiddleware(http.HandlerFunc(userController.UpdateUser)))

	userMux.Handle("/login", http.HandlerFunc(handlers.Login))
	userMux.Handle("/register", http.HandlerFunc(handlers.Register))
	userMux.Handle("/checkCode", http.HandlerFunc(handlers.CheckCode))
	return userMux
}
