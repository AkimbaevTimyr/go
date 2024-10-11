package repository

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
)

type CategoryRepository interface {
	Create(request models.Category) (*models.Category, *helpers.Error)
}

type categoryRepo struct {
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepo{}
}

func (r *categoryRepo) Create(model models.Category) (*models.Category, *helpers.Error) {
	if err := database.DB.Create(&model).Error; err != nil {
		return nil, &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
	}
	return &model, nil
}
