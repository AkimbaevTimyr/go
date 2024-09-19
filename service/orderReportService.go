package service

import (
	"akimbaev/database"
	"akimbaev/models"
	"akimbaev/repository"
	"fmt"
	"log"
)

type OrderReportService interface {
	Connect(id int, userId uint) (*models.OrderReport, error)
}

type orderReportService struct {
	repo repository.OrderReportRepository
}

func NewOrderReportService(repo repository.OrderReportRepository) OrderReportService {
	return &orderReportService{
		repo: repo,
	}
}

func (repo *orderReportService) Connect(id int, userId uint) (*models.OrderReport, error) {
	// tx := database.DB.Begin()

	var order models.Order

	result := database.DB.First(&models.Order{}, id).Error

	if result != nil {
		return nil, fmt.Errorf("order with id %d not found", id)
	}

	//создание отчета
	Report := models.OrderReport{
		UserId:  userId,
		OrderId: uint(id),
	}

	if err := database.DB.Create(&Report).Error; err != nil {
		// tx.Rollback()
		log.Fatalln(err)
		return nil, err
	}

	database.DB.First(&order, id)
	database.DB.Preload("User").First(&order, order.ID)

	order.User.Balance -= order.Price
	if err := database.DB.Save(&order.User).Error; err != nil {
		log.Fatalln(err, "123")
		// tx.Rollback()
		return nil, err
	}

	database.DB.Preload("Order").Find(&Report, Report.ID)

	// tx.Commit()
	return &Report, nil
}
