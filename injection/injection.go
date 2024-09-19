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
