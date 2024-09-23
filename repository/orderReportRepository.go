package repository

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
)

type OrderReportRepository interface {
	Create(orderId, userId uint) (*models.OrderReport, *helpers.Error)
}

type orderReportRepository struct{}

func NewOrderReportRepository() OrderReportRepository {
	return &orderReportRepository{}
}

func (r *orderReportRepository) Create(orderId, userId uint) (*models.OrderReport, *helpers.Error) {

	report := models.OrderReport{
		UserId:  userId,
		OrderId: orderId,
	}

	if err := database.DB.Create(&report).Error; err != nil {
		return nil, &helpers.Error{Code: helpers.EINTERNAL, Message: "internal server error"}
	}

	return &report, nil
}
