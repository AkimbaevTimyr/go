package service

import (
	"akimbaev/models"
	"akimbaev/repository"
	"akimbaev/requests"
	"akimbaev/requests/order"
)

type OrderService interface {
	GetOrders(id int, params order.IndexRequest) (*[]models.Order, error)
	CreateOrder(id int, request requests.OrderRequest) (*models.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (s *orderService) GetOrders(id int, params order.IndexRequest) (*[]models.Order, error) {

	orders, err := s.repo.GetUserOrders(id, params)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *orderService) CreateOrder(id int, request requests.OrderRequest) (*models.Order, error) {

	order, err := s.repo.CreateOrder(id, request)

	if err != nil {
		return nil, err
	}

	return order, nil

}
