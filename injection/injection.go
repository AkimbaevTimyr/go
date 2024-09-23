package injection

import (
	"akimbaev/controller"
	"akimbaev/repository"
	"akimbaev/service"
)

func InitUserController() *controller.UserController {
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	return userController
}

func InitOrderReportController() *controller.UserReportController {
	reportRepository := repository.NewOrderReportRepository()
	orderRepository := repository.NewOrderRepository()
	service := service.NewOrderReportService(reportRepository, orderRepository)
	controller := controller.NewUserReportController(service)
	return controller
}

func InitAuthController() *controller.AuthController {
	service := service.NewAuthService()
	controller := controller.NewAuthController(service)
	return controller
}

func InitOrderController() *controller.OrderController {
	repository := repository.NewOrderRepository()
	service := service.NewOrderService(repository)
	controller := controller.NewOrderController(service)
	return controller
}
