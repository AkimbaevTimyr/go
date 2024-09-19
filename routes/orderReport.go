package routes

import (
	"akimbaev/injection"
	"akimbaev/middleware"
	"net/http"
)

func OrderReportMux() http.Handler {
	orderReportMux := http.NewServeMux()

	orderReportController := injection.InitOrderReportController()
	orderReportMux.Handle("/connect", middleware.AuthMiddleware(http.HandlerFunc(orderReportController.Connect)))

	return orderReportMux
}
