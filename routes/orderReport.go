package routes

import (
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func OrderReportMux() http.Handler {
	mux := http.NewServeMux()

	c := injection.InitOrderReportController()

	mux.Handle("/connect", middleware.AuthMiddleware(http.HandlerFunc(c.Connect)))
	mux.Handle("/myReports", middleware.CreateMiddleware(middleware.AuthMiddleware)(http.HandlerFunc(c.MyReports)))
	mux.Handle("/show", middleware.CreateMiddleware(middleware.AuthMiddleware)(http.HandlerFunc(c.Show)))

	return mux
}
