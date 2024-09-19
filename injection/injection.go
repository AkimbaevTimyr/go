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
	repository := repository.NewOrderReportRepository()
	service := service.NewOrderReportService(repository)
	controller := controller.NewUserReportController(service)
	return controller
}
