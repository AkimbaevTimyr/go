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

func InitAuthController() *controller.AuthController {
	service := service.NewAuthService()
	controller := controller.NewAuthController(service)
	return controller
}

func InitOrderController() *controller.PostController {
	repository := repository.NewPostRepository()
	service := service.NewPostService(repository)
	controller := controller.NewPostController(service)
	return controller
}

func InitSubscriptionController() *controller.SubscriptionController {
	subsRepo := repository.NewSubscriptionRepository()
	planRepo := repository.NewPlanRepository()
	userRepo := repository.NewUserRepository()
	service := service.NewSubscriptionService(subsRepo, planRepo, userRepo)
	controller := controller.NewSubscriptionController(service)
	return controller
}

func InitCategoryController() *controller.Category {
	r := repository.NewCategoryRepository()
	s := service.NewCategoryService(r)
	c := controller.NewCategoryController(s)
	return c
}
