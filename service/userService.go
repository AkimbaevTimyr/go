package service

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/repository"
	"akimbaev/requests"
)

type UserService interface {
	GetUser(id int) (*models.User, *helpers.Error)
	DeleteUser(id int) *helpers.Error
	UpdateUser(id int, request requests.UpdateUserRequest) (*models.User, *helpers.Error)
	AddBalance(id int, request requests.AddBalanceRequest) *helpers.Error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUser(id int) (*models.User, *helpers.Error) {
	user, err := s.repo.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id int) *helpers.Error {
	err := s.repo.DeleteUserById(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUser(id int, request requests.UpdateUserRequest) (*models.User, *helpers.Error) {
	user, err := s.repo.UpdateUser(id, request)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) AddBalance(id int, request requests.AddBalanceRequest) *helpers.Error {
	//find user by id
	user, err := s.repo.GetUserById(id)

	if err != nil {
		return err
	}

	//update user balance
	user.Balance += request.Amount

	//save user
	database.DB.Save(&user)

	return nil
}
