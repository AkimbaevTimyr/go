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
	GetById(id int) (*models.Order, error)
	DeleteById(id int) error
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

func (r *orderRepo) GetById(id int) (*models.Order, error) {
	var order models.Order

	if err := database.DB.First(&order, id).Error; err != nil {
		return nil, fmt.Errorf("order with id %d not found", id)
	}
	return &order, nil
}

func (r *orderRepo) DeleteById(id int) error {
	if err := database.DB.Delete(&models.Order{}, id).Error; err != nil {
		return fmt.Errorf("order with id %d not found", id)
	}

	result := database.DB.Delete(&models.Order{}, id)

	if result.Error != nil {
		return fmt.Errorf("internal server error")
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("order with id %d not found", id)
	}
	return nil
}
