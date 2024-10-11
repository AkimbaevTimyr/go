package routes

import (
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func CategoryMux() http.Handler {
	mux := http.NewServeMux()

	c := injection.InitCategoryController()

	mux.Handle("/create", middleware.AuthMiddleware(http.HandlerFunc(c.Create)))

	return mux
}
