package service

import (
	"akimbaev/models"
	"akimbaev/repository"
	"akimbaev/requests"
)

type UserService interface {
	GetUser(id int) (*models.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, request requests.UpdateUserRequest) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetUser(id int) (*models.User, error) {
	user, err := s.repo.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id int) error {
	err := s.repo.DeleteUserById(id)

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateUser(id int, request requests.UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.UpdateUser(id, request)

	if err != nil {
		return nil, err
	}

	return user, nil
}
