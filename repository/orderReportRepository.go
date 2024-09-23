package repository

import (
	"akimbaev/database"
	"akimbaev/models"
	"fmt"
)

type OrderReportRepository interface {
	Create(orderId, userId uint) (*models.OrderReport, error)
}

type orderReportRepository struct{}

func NewOrderReportRepository() OrderReportRepository {
	return &orderReportRepository{}
}

func (r *orderReportRepository) Create(orderId, userId uint) (*models.OrderReport, error) {

	report := models.OrderReport{
		UserId:  userId,
		OrderId: orderId,
	}

	if err := database.DB.Create(&report).Error; err != nil {
		return nil, fmt.Errorf("cannot create orderReport")
	}

	return &report, nil
}
