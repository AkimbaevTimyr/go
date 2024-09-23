package service

import (
	"akimbaev/database"
	"akimbaev/models"
	"akimbaev/repository"
)

type OrderReportService interface {
	Connect(id int, userId uint) (*models.OrderReport, error)
}

type orderReportService struct {
	repo      repository.OrderReportRepository
	orderRepo repository.OrderRepository
}

func NewOrderReportService(repo repository.OrderReportRepository, orderRepo repository.OrderRepository) OrderReportService {
	return &orderReportService{
		repo:      repo,
		orderRepo: orderRepo,
	}
}

// проверить
func (s *orderReportService) Connect(id int, userId uint) (*models.OrderReport, error) {

	order, err := s.orderRepo.GetById(id)

	if err != nil {
		return nil, err
	}

	report, err := s.repo.Create(uint(id), userId)

	if err != nil {
		return nil, err
	}

	database.DB.First(&order, id)
	database.DB.Preload("User").First(&order, order.ID)

	order.User.Balance -= order.Price
	if err := database.DB.Save(&order.User).Error; err != nil {
		return nil, err
	}

	database.DB.Preload("Order").Find(&report, report.ID)

	return report, nil
}
