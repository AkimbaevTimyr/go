package repository

import (
	"akimbaev/database"
	"akimbaev/models"
	"akimbaev/requests"
	"akimbaev/requests/order"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetUserOrders(id int, params order.IndexRequest) (*[]models.Order, error)
	CreateOrder(id int, request requests.OrderRequest) (*models.Order, error)
}

type orderRepo struct{}

func NewOrderRepository() OrderRepository {
	return &orderRepo{}
}

func (r *orderRepo) GetUserOrders(id int, params order.IndexRequest) (*[]models.Order, error) {

	var orders []models.Order

	database.DB.Scopes(func(db *gorm.DB) *gorm.DB {
		return database.Paginate(db, params.Page, params.Sort, params.Count)
	}).Where("user_id = ?", id).Find(&orders)

	return &orders, nil
}

func (r *orderRepo) CreateOrder(id int, request requests.OrderRequest) (*models.Order, error) {

	NewOrder := models.Order{
		Title:   request.Title,
		Content: request.Content,
		Price:   request.Price,
		UserId:  uint(id),
	}

	if err := database.DB.Create(&NewOrder).Error; err != nil {
		return nil, fmt.Errorf("something where wrong: %v", err.Error())
	}

	return &NewOrder, nil
}
