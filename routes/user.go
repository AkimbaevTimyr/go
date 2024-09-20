package routes

import (
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

	return userMux
}
