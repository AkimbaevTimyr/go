package repository

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/requests"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserById(id int) (*models.User, *helpers.Error)
	DeleteUserById(id int) *helpers.Error
	UpdateUser(id int, request requests.UpdateUserRequest) (*models.User, *helpers.Error)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) GetUserById(id int) (*models.User, *helpers.Error) {
	var user models.User

	err := database.DB.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helpers.Error{Code: helpers.ENOTFOUND, Message: fmt.Sprintf("user with id %d not found", id)}
		}
		return nil, &helpers.Error{Code: helpers.ENOTFOUND, Message: "internal server error"}
	}

	if err := database.DB.Preload("Orders").First(&user, user.ID).Error; err != nil {
		return nil, &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
	}

	return &user, nil
}

func (r *userRepo) DeleteUserById(id int) *helpers.Error {
	result := database.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		return &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
	}

	if result.RowsAffected == 0 {
		return &helpers.Error{Code: helpers.ENOTFOUND, Message: fmt.Sprintf("user with id %d not found", id)}
	}

	return nil
}

func (r *userRepo) UpdateUser(id int, request requests.UpdateUserRequest) (*models.User, *helpers.Error) {
	var user models.User

	err := database.DB.First(&user, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helpers.Error{Code: helpers.ENOTFOUND, Message: fmt.Sprintf("user with id %d not found", id)}
		}
		return nil, &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
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
