package repository

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type PlanRepository interface {
	GetById(id int) (*models.Plan, *helpers.Error)
}

type planRepo struct {
}

func NewPlanRepository() PlanRepository {
	return &planRepo{}
}

func (r *planRepo) GetById(id int) (*models.Plan, *helpers.Error) {
	var plan models.Plan

	err := database.DB.Where("id = ?", id).First(&plan).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &helpers.Error{Code: helpers.ENOTFOUND, Message: fmt.Sprintf("plan with id %d not found", id)}
		}
		return nil, &helpers.Error{Code: helpers.ENOTFOUND, Message: "internal server error"}
	}

	return &plan, nil
}
