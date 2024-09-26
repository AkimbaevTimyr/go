package routes

import (
	"akimbaev/injection"
	"net/http"
)

func AuthMux() http.Handler {
	mux := http.NewServeMux()

	c := injection.InitAuthController()

	mux.Handle("/login", http.HandlerFunc(c.Login))
	mux.Handle("/register", http.HandlerFunc(c.Register))
	mux.Handle("/checkCode", http.HandlerFunc(c.CheckCode))

	return mux
}
