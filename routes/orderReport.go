package routes

import (
	"akimbaev/handlers"
	"akimbaev/middleware"
	"net/http"
)

func OrderReportMux() http.Handler {
	orderReportMux := http.NewServeMux()
	orderReportMux.Handle("/connect", middleware.AuthMiddleware(http.HandlerFunc(handlers.Connect)))

	return orderReportMux
}
