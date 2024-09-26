package service

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/repository"
	report "akimbaev/requests/reports"
)

type OrderReportService interface {
	Connect(id int, userId uint) (*models.OrderReport, *helpers.Error)
	MyReports(userId uint, params report.IndexRequest) (*[]models.OrderReport, *helpers.Error)
	Show(id int) (*models.OrderReport, *helpers.Error)
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
func (s *orderReportService) Connect(id int, userId uint) (*models.OrderReport, *helpers.Error) {

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
	database.DB.Save(&order.User)

	database.DB.Preload("Order").Find(&report, report.ID)

	return report, nil
}

func (s *orderReportService) MyReports(userId uint, params report.IndexRequest) (*[]models.OrderReport, *helpers.Error) {
	orders, err := s.repo.GetReportsByUserId(userId, params)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderReportService) Show(id int) (*models.OrderReport, *helpers.Error) {
	order, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return order, nil
}
