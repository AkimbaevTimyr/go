package routes

import (
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func UserMux() http.Handler {
	mux := http.NewServeMux()

	c := injection.InitUserController()

	mux.Handle("/get", middleware.AuthMiddleware(http.HandlerFunc(c.GetUser)))
	mux.Handle("/delete", middleware.AuthMiddleware(http.HandlerFunc(c.DeleteUser)))
	mux.Handle("/update", middleware.AuthMiddleware(http.HandlerFunc(c.UpdateUser)))

	return mux
}
