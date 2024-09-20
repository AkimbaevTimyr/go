package routes

import (
	"akimbaev/injection"
	"net/http"
)

func AuthMux() http.Handler {
	authMux := http.NewServeMux()

	authController := injection.InitAuthController()

	authMux.Handle("/login", http.HandlerFunc(authController.Login))
	authMux.Handle("/register", http.HandlerFunc(authController.Register))
	authMux.Handle("/checkCode", http.HandlerFunc(authController.CheckCode))
	return authMux
}
