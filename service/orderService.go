package service

import (
	"akimbaev/database"
	"akimbaev/helpers"
	"akimbaev/models"
	"akimbaev/repository"
	"akimbaev/requests"
	"akimbaev/requests/order"
)

type OrderService interface {
	GetOrders(id int, params order.IndexRequest) (*[]models.Order, *helpers.Error)
	CreateOrder(id int, request requests.OrderRequest) (*models.Order, *helpers.Error)
	ChangeStatus(id int, status string) *helpers.Error
	Delete(id int) *helpers.Error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (s *orderService) GetOrders(id int, params order.IndexRequest) (*[]models.Order, *helpers.Error) {

	orders, err := s.repo.GetUserOrders(id, params)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderService) CreateOrder(id int, request requests.OrderRequest) (*models.Order, *helpers.Error) {

	order, err := s.repo.CreateOrder(id, request)

	if err != nil {
		return nil, err
	}

	return order, nil

}

func (s *orderService) ChangeStatus(id int, status string) *helpers.Error {
	order, err := s.repo.GetById(id)

	if err != nil {
		return err
	}

	order.Status = models.StatusMap[status]

	if err := database.DB.Save(&order).Error; err != nil {
		return nil
	}

	return nil
}

func (s *orderService) Delete(id int) *helpers.Error {

	err := s.repo.DeleteById(id)

	if err != nil {
		return err
	}

	return nil

}
