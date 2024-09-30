package routes

import (
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func SubscriptionMux() http.Handler {
	mux := http.NewServeMux()

	c := injection.InitSubscriptionController()

	mux.Handle("/purchase", middleware.AuthMiddleware(http.HandlerFunc(c.Purchase)))

	return mux
}
