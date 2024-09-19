package repository

import (
	"akimbaev/database"
	"akimbaev/models"
	"akimbaev/requests"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(id int) (*models.User, error)
	DeleteUserById(id int) error
	UpdateUser(id int, request requests.UpdateUserRequest) (*models.User, error)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) GetUserById(id int) (*models.User, error) {
	var user models.User

	err := database.DB.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}

	if err := database.DB.Preload("Orders").First(&user, user.ID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) DeleteUserById(id int) error {
	result := database.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		return fmt.Errorf("something where wrong: %v", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}

func (r *userRepo) UpdateUser(id int, request requests.UpdateUserRequest) (*models.User, error) {
	var user models.User

	err := database.DB.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, fmt.Errorf("something where wrong: %v", err.Error())
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	database.DB.Save(&user)

	return &user, nil
}
